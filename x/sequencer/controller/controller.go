package controller

import (
	"math"
	"os"
	"path"
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

	store  *Store
	config *config.Config
}

func NewController(log log.Logger, homeDir string) (out ArweaveBlocksController) {
	controller := new(SyncerController)
	InitLogger(log, logrus.InfoLevel.String())

	var err error
	filepath := path.Join(homeDir, "syncer.json")
	if _, err := os.Stat(filepath); err != nil {
		// Empty file path loads default config
		filepath = ""
	}

	controller.config, err = config.Load(filepath)
	if err != nil {
		panic(err)
	}

	out = controller
	return
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
	controller.Task = task.NewTask(controller.config, "controller")

	if !controller.config.Syncer.Enabled {
		return
	}

	// FIXME: Add monitor to prometheus
	monitor := monitor_syncer.NewMonitor().
		WithMaxHistorySize(30)

	watched := func() *task.Task {
		client := arweave.NewClient(controller.Ctx, controller.config).
			WithTagValidator(warp.ValidateTag)

		monitor := monitor_syncer.NewMonitor().
			WithMaxHistorySize(30)

		networkMonitor := listener.NewNetworkMonitor(controller.config).
			WithClient(client).
			WithMonitor(monitor).
			WithInterval(controller.config.NetworkMonitor.Period).
			WithRequiredConfirmationBlocks(20)

		blockDownloader := listener.NewBlockDownloader(controller.config).
			WithClient(client).
			WithInputChannel(networkMonitor.Output).
			WithMonitor(monitor).
			WithBackoff(0, controller.config.Syncer.TransactionMaxInterval).
			WithHeightRange(initHeight, math.MaxUint64)

		transactionDownloader := listener.NewTransactionDownloader(controller.config).
			WithClient(client).
			WithInputChannel(blockDownloader.Output).
			WithMonitor(monitor).
			WithBackoff(0, controller.config.Syncer.TransactionMaxInterval).
			WithFilterInteractions()

		store := NewStore(controller.config).
			WithInputChannel(transactionDownloader.Output).
			WithMonitor(monitor)
		controller.store = store

		return task.NewTask(controller.config, "watched").
			WithSubtask(networkMonitor.Task).
			WithSubtask(blockDownloader.Task).
			WithSubtask(transactionDownloader.Task).
			WithSubtask(store.Task)
	}

	watchdog := task.NewWatchdog(controller.config).
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
