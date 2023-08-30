package arweave

import (
	"sync"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	syncer_arweave "github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/config"
	"github.com/warp-contracts/syncer/src/utils/listener"
	"github.com/warp-contracts/syncer/src/utils/monitoring"
	"github.com/warp-contracts/syncer/src/utils/task"
)

type Store struct {
	*task.Task

	monitor monitoring.Monitor
	mtx     sync.Mutex

	input  chan *listener.Payload
	blocks []types.NextArweaveBlock
}

func NewStore(config *config.Config) (store *Store) {
	store = new(Store)

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
			Hash:      payload.BlockHash,
		},
		Transactions: transactions(payload),
	}

	store.mtx.Lock()
	store.blocks = append(store.blocks, block)
	store.mtx.Unlock()
}

func transactions(payload *listener.Payload) []syncer_arweave.Transaction {
	txs := make([]syncer_arweave.Transaction, 0, len(payload.Transactions))
	for _, tx := range payload.Transactions {
		txs = append(txs, *tx)
	}
	return txs
}

func (store *Store) getAndClearBlocks() []types.NextArweaveBlock {
	store.mtx.Lock()
	defer store.mtx.Unlock()

	blocks := store.blocks
	store.blocks = nil

	return blocks
}
