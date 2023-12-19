package controller

import (
	"sort"
	"sync"

	"github.com/gammazero/deque"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	syncer_arweave "github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/config"
	"github.com/warp-contracts/syncer/src/utils/listener"
	"github.com/warp-contracts/syncer/src/utils/monitoring"
	"github.com/warp-contracts/syncer/src/utils/smartweave"
	"github.com/warp-contracts/syncer/src/utils/task"
	"github.com/warp-contracts/syncer/src/utils/warp"
)

type Store struct {
	*task.Task

	monitor monitoring.Monitor
	mtx     sync.RWMutex

	input  chan *listener.Payload
	blocks *deque.Deque[*types.NextArweaveBlock]
}

func NewStore(config *config.Config) (self *Store) {
	self = new(Store)
	self.blocks = deque.New[*types.NextArweaveBlock]()

	self.Task = task.NewTask(config, "store").
		WithSubtaskFunc(self.run)

	return
}

func (self *Store) WithMonitor(v monitoring.Monitor) *Store {
	self.monitor = v
	return self
}

func (self *Store) WithInputChannel(v chan *listener.Payload) *Store {
	self.input = v
	return self
}

func (self *Store) run() (err error) {
	for {
		select {
		case payload := <-self.input:
			self.processPayload(payload)
		case <-self.Ctx.Done():
			return
		}
	}
}

func (self *Store) processPayload(payload *listener.Payload) {
	self.Log.Debug("-> processPayload")
	defer self.Log.Debug("<- processPayload")

	block := types.NextArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    uint64(payload.BlockHeight),
			Timestamp: uint64(payload.BlockTimestamp),
			Hash:      payload.BlockHash.Base64(),
		},
		Transactions: self.transactions(payload),
	}

	self.mtx.Lock()
	// TODO add saving to database
	self.blocks.PushBack(&block)
	self.Log.
		WithField("queue_size", self.blocks.Len()).
		WithField("height", block.BlockInfo.Height).
		Info("Arweave block added to the local queue")
	self.mtx.Unlock()
}

func (self *Store) transactions(payload *listener.Payload) []*types.ArweaveTransaction {
	self.Log.Debug("-> transactions")
	defer self.Log.Debug("<- transactions")

	txs := make([]*types.ArweaveTransaction, 0, len(payload.Transactions))
	for _, tx := range payload.Transactions {
		contract := getContractFromTag(tx)
		if len(contract) == 0 {
			self.Log.WithField("tx_id", tx.ID.Base64()).Warn("Ignored interaction without a contract tag")
			continue
		}

		txs = append(txs, &types.ArweaveTransaction{
			Id:       tx.ID.Base64(),
			Contract: contract,
			SortKey:  warp.CreateSortKey(tx.ID, payload.BlockHeight, payload.BlockHash),
		})
	}
	// sort transactions by sort key
	sort.Slice(txs, func(i, j int) bool {
		return txs[i].SortKey < txs[j].SortKey
	})
	return txs
}

func getContractFromTag(tx *syncer_arweave.Transaction) string {
	for _, tag := range tx.Tags {
		if string(tag.Name) == smartweave.TagContractTxId {
			return string(tag.Value)
		}
	}
	return ""
}

func (self *Store) GetNextArweaveBlock(height uint64) *types.NextArweaveBlock {
	self.Log.Debug("-> SetLastAcceptedBlock")
	defer self.Log.Debug("<- SetLastAcceptedBlock")

	self.mtx.RLock()
	defer self.mtx.RUnlock()

	for i := 0; i < self.blocks.Len(); i++ {
		block := self.blocks.At(i)
		if block.BlockInfo.Height == height {
			return block
		}
	}
	return nil
}

func (self *Store) RemoveNextArweaveBlocksUpToHeight(height uint64) {
	self.Log.Debug("-> SetLastAcceptedBlock")
	defer self.Log.Debug("<- SetLastAcceptedBlock")

	self.mtx.Lock()
	defer self.mtx.Unlock()

	for self.blocks.Len() > 0 {
		block := self.blocks.Front()
		if block.BlockInfo.Height <= height {
			self.blocks.PopFront()
			self.Log.
				WithField("queue_size", self.blocks.Len()).
				WithField("height", block.BlockInfo.Height).
				Info("Arweave block removed from the local queue")
		} else {
			break
		}
	}
}
