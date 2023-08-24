package arweave

import (
	syncerarweave "github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/config"
	"github.com/warp-contracts/syncer/src/utils/listener"
	"github.com/warp-contracts/syncer/src/utils/monitoring"
	"github.com/warp-contracts/syncer/src/utils/task"
)

type Store struct {
	*task.Processor[*listener.Payload, *syncerarweave.Transaction]

	monitor monitoring.Monitor
}

func NewStore(config *config.Config) (self *Store) {
	self = new(Store)

	self.Processor = task.NewProcessor[*listener.Payload, *syncerarweave.Transaction](config, "store").
		WithBatchSize(config.Syncer.StoreBatchSize).
		WithOnFlush(config.Syncer.StoreMaxTimeInQueue, self.flush).
		WithOnProcess(self.process).
		WithBackoff(0, config.Syncer.StoreMaxBackoffInterval)

	return
}

func (store *Store) WithMonitor(v monitoring.Monitor) *Store {
	store.monitor = v
	return store
}

func (store *Store) WithInputChannel(v chan *listener.Payload) *Store {
	store.Processor = store.Processor.WithInputChannel(v)
	return store
}

func (store *Store) flush(data []*syncerarweave.Transaction) (out []*syncerarweave.Transaction, err error) {
	return data, nil
}

func (store *Store) process(payload *listener.Payload) (out []*syncerarweave.Transaction, err error) {
	out = payload.Transactions
	// TODO save in store
	store.Log.
		WithField("BlockHeight", payload.BlockHeight).
		WithField("BlockTimestamp", payload.BlockTimestamp).
		Info("Arweave block")
	return
}
