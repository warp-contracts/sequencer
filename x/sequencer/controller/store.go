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

func NewStore(config *config.Config) (store *Store) {
	store = new(Store)
	store.blocks = deque.New[*types.NextArweaveBlock]()

	store.Task = task.NewTask(config, "store").
		WithSubtaskFunc(store.run)

	return
}

func (store *Store) WithMonitor(v monitoring.Monitor) *Store {
	store.monitor = v
	return store
}

func (store *Store) WithInputChannel(v chan *listener.Payload) *Store {
	store.input = v
	return store
}

func (store *Store) run() (err error) {
	for {
		select {
		case payload := <-store.input:
			store.processPayload(payload)
		case <-store.Ctx.Done():
			return
		}
	}
}

func (store *Store) processPayload(payload *listener.Payload) {
	block := types.NextArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    uint64(payload.BlockHeight),
			Timestamp: uint64(payload.BlockTimestamp),
			Hash:      payload.BlockHash.Base64(),
		},
		Transactions: store.transactions(payload),
	}

	store.mtx.Lock()
	// TODO add saving to database
	store.blocks.PushBack(&block)
	store.mtx.Unlock()
}

func (store *Store) transactions(payload *listener.Payload) []*types.ArweaveTransaction {
	txs := make([]*types.ArweaveTransaction, 0, len(payload.Transactions))
	for _, tx := range payload.Transactions {
		err := warp.ValidateTags(tx.Tags)
		if err != nil {
			store.Log.WithError(err).WithField("tx_id", tx.ID.Base64()).Warn("Ignored interaction with invalid tags")
			continue
		}

		contract := getContractFromTag(tx)
		if len(contract) == 0 {
			store.Log.WithField("tx_id", tx.ID.Base64()).Warn("Ignored interaction without a contract tag")
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

func (store *Store) GetNextArweaveBlock(height uint64) *types.NextArweaveBlock {
	store.mtx.RLock()
	defer store.mtx.RUnlock()

	for i := 0; i < store.blocks.Len(); i++ {
		block := store.blocks.At(i)
		if block.BlockInfo.Height == height {
			return block
		}
	}
	return nil
}

func (store *Store) RemoveNextArweaveBlocksUpToHeight(height uint64) {
	store.mtx.Lock()
	defer store.mtx.Unlock()

	for store.blocks.Len() > 0 && store.blocks.Front().BlockInfo.Height <= height {
		store.blocks.PopFront()
	}
}
