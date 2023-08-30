package arweave

import (
	"math"
	"sync/atomic"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"

	"github.com/warp-contracts/syncer/src/sync"
	"github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/config"
	"github.com/warp-contracts/syncer/src/utils/listener"
	monitor_syncer "github.com/warp-contracts/syncer/src/utils/monitoring/syncer"
	"github.com/warp-contracts/syncer/src/utils/task"
	"github.com/warp-contracts/syncer/src/utils/warp"
)

type ArweaveBlocksController struct {
	sync.Controller

	store  *Store
	keeper keeper.Keeper
	IsRunning *atomic.Bool
}

func CreateController(keeper keeper.Keeper) *ArweaveBlocksController {
	controller := new(ArweaveBlocksController)
	controller.keeper = keeper
	controller.IsRunning = &atomic.Bool{}

	return controller
}

// TODO add controller stop
func (controller *ArweaveBlocksController) StartController(initHeight uint64) {
	controller.initController(initHeight)

	err := controller.Start()
	if err != nil {
		panic(err)
	}
	controller.IsRunning.Store(true)
}

func (controller *ArweaveBlocksController) initController(initHeight uint64) {
	var config = config.Default()

	controller.Task = task.NewTask(config, "controller")

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
			WithRequiredConfirmationBlocks(25)

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

func (controller *ArweaveBlocksController) StoreArweaveBlocks(ctx sdk.Context) {
	if controller.IsRunning.Load() {
		for _, block := range controller.store.getAndClearBlocks() {
			controller.keeper.SetNextArweaveBlock(ctx, block)
		}	
	}
}
