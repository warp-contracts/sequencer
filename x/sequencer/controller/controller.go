package controller

import (
	"math"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/sirupsen/logrus"
	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/sync"
	"github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/config"
	"github.com/warp-contracts/syncer/src/utils/listener"
	monitor_syncer "github.com/warp-contracts/syncer/src/utils/monitoring/syncer"
	"github.com/warp-contracts/syncer/src/utils/task"
	"github.com/warp-contracts/syncer/src/utils/warp"
)

// Controller for fetching Arweave blocks to add them to the sequencer blockchain or validate blocks added by the Proposer.
type ArweaveBlocksController interface {
	// Starts the fetching of Arweave blocks beginning from the given height
	Start(initHeight uint64)

	// Has the controller been started?
	IsRunning() bool

	// Returns the fetched Arweave block with the given height
	GetNextArweaveBlock(height uint64) *types.NextArweaveBlock

	// Deletes all fetched Arweave blocks with height not greater than the given one
	RemoveNextArweaveBlocksUpToHeight(height uint64)
}

type SyncerController struct {
	sync.Controller

	store *Store
}

func NewController(log log.Logger) ArweaveBlocksController {
	controller := new(SyncerController)
	InitLogger(log, logrus.InfoLevel.String())
	return controller
}

// TODO add controller stop
func (controller *SyncerController) Start(initHeight uint64) {
	controller.initController(initHeight)

	err := controller.Controller.Start()
	if err != nil {
		panic(err)
	}
}

func (controller *SyncerController) initController(initHeight uint64) {
	// TODO give option to override default values
	var config = config.Default()

	controller.Task = task.NewTask(config, "controller")

	// FIXME: Add monitor to prometheus
	monitor := monitor_syncer.NewMonitor().
		WithMaxHistorySize(30)

	watched := func() *task.Task {
		client := arweave.NewClient(controller.Ctx, config).
			WithTagValidator(warp.ValidateTag)

		monitor := monitor_syncer.NewMonitor().
			WithMaxHistorySize(30)

		networkMonitor := listener.NewNetworkMonitor(config).
			WithClient(client).
			WithMonitor(monitor).
			WithInterval(config.NetworkMonitor.Period).
			WithRequiredConfirmationBlocks(20)

		blockDownloader := listener.NewBlockDownloader(config).
			WithClient(client).
			WithInputChannel(networkMonitor.Output).
			WithMonitor(monitor).
			WithBackoff(0, config.Syncer.TransactionMaxInterval).
			WithHeightRange(initHeight, math.MaxUint64)

		transactionDownloader := listener.NewTransactionDownloader(config).
			WithClient(client).
			WithInputChannel(blockDownloader.Output).
			WithMonitor(monitor).
			WithBackoff(0, config.Syncer.TransactionMaxInterval).
			WithFilterInteractions()

		store := NewStore(config).
			WithInputChannel(transactionDownloader.Output).
			WithMonitor(monitor)
		controller.store = store

		return task.NewTask(config, "watched").
			WithSubtask(networkMonitor.Task).
			WithSubtask(blockDownloader.Task).
			WithSubtask(transactionDownloader.Task).
			WithSubtask(store.Task)
	}

	watchdog := task.NewWatchdog(config).
		WithTask(watched).
		WithIsOK(30*time.Second, func() bool {
			isOK := monitor.IsOK()
			if !isOK {
				monitor.Clear()
				monitor.GetReport().Run.Errors.NumWatchdogRestarts.Inc()
			}
			return isOK
		})

	controller.Task = controller.Task.
		WithSubtask(monitor.Task).
		WithSubtask(watchdog.Task)
}

func (controller *SyncerController) IsRunning() bool {
	return controller.Controller.Task != nil && !controller.Controller.Task.IsStopping.Load()
}

func (controller *SyncerController) GetNextArweaveBlock(height uint64) *types.NextArweaveBlock {
	if controller.IsRunning() {
		return controller.store.GetNextArweaveBlock(height)
	}
	return nil
}

func (controller *SyncerController) RemoveNextArweaveBlocksUpToHeight(height uint64) {
	if controller.IsRunning() {
		controller.store.removeNextArweaveBlocksUpToHeight(height)
	}
}
