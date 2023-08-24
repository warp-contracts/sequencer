package arweave

import (
	"time"

	"github.com/warp-contracts/syncer/src/sync"
	"github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/config"
	"github.com/warp-contracts/syncer/src/utils/listener"
	monitor_syncer "github.com/warp-contracts/syncer/src/utils/monitoring/syncer"
	"github.com/warp-contracts/syncer/src/utils/task"
	"github.com/warp-contracts/syncer/src/utils/warp"
)

func Controller() {
	controller, err := newController()
	if err != nil {
		panic(err)
	}

	err = controller.Start()
	if err != nil {
		panic(err)
	}

	// FIXME
	<-controller.CtxRunning.Done()

	controller.StopWait()
}

func newController() (controller *sync.Controller, err error) {
	var config = config.Default()

	controller = new(sync.Controller)
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
			WithRequiredConfirmationBlocks(config.NetworkMonitor.RequiredConfirmationBlocks)

		blockDownloader := listener.NewBlockDownloader(config).
			WithClient(client).
			WithInputChannel(networkMonitor.Output).
			WithMonitor(monitor).
			WithBackoff(0, config.Syncer.TransactionMaxInterval).
			WithHeightRange(1246586, 1256586) // FIXME

		transactionDownloader := listener.NewTransactionDownloader(config).
			WithClient(client).
			WithInputChannel(blockDownloader.Output).
			WithMonitor(monitor).
			WithBackoff(0, config.Syncer.TransactionMaxInterval).
			WithFilterInteractions()

		store := NewStore(config).
			WithInputChannel(transactionDownloader.Output).
			WithMonitor(monitor)

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

	return
}
