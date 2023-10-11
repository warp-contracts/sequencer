package controller

import (
	"math"
	"os"
	"path"
	"sync"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/sirupsen/logrus"
	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/config"
	"github.com/warp-contracts/syncer/src/utils/listener"
	"github.com/warp-contracts/syncer/src/utils/monitoring"
	monitor_syncer "github.com/warp-contracts/syncer/src/utils/monitoring/syncer"
	"github.com/warp-contracts/syncer/src/utils/task"
	"github.com/warp-contracts/syncer/src/utils/warp"
)

// Controller for fetching Arweave blocks to add them to the sequencer blockchain or validate blocks added by the Proposer.
type ArweaveBlocksController interface {
	// Sets the last block height accepted by the sequencer network
	SetLastAcceptedBlockHeight(uint64)

	// Gracefully stops the controller, waits for all tasks to finish
	StopWait()

	// Returns the fetched Arweave block with the given height
	GetNextArweaveBlock(height uint64) *types.NextArweaveBlock
}

type SyncerController struct {
	*task.Task

	store  *Store
	config *config.Config

	// Runtime state
	mtx                       sync.Mutex
	lastAcceptedArweaveHeight uint64
	blockDownloader           *listener.BlockDownloader
}

func NewController(log log.Logger, configPath string) (out ArweaveBlocksController, err error) {
	self := new(SyncerController)
	InitLogger(log, logrus.InfoLevel.String())

	// Load configuration from path, env or defaults
	filepath := path.Join(configPath, "syncer.json")
	if _, err := os.Stat(filepath); err != nil {
		// Empty file path loads default config
		filepath = ""
	}

	self.config, err = config.Load(filepath)
	if err != nil {
		return
	}

	// Setup the tasks
	self.Task = task.NewTask(self.config, "controller")

	// Monitoring and performance metrics
	monitor := monitor_syncer.NewMonitor().
		WithMaxHistorySize(30)

	server := monitoring.NewServer(self.config).
		WithMonitor(monitor)

	// Function that creates the watched task
	// This can be called multiple times to setup syncing when something got stuck
	watched := func() *task.Task {
		self.mtx.Lock()
		defer self.mtx.Unlock()

		client := arweave.NewClient(self.Ctx, self.config).
			WithTagValidator(warp.ValidateTag)

		networkMonitor := listener.NewNetworkMonitor(self.config).
			WithClient(client).
			WithMonitor(monitor).
			WithInterval(self.config.NetworkMonitor.Period).
			WithRequiredConfirmationBlocks(self.config.NetworkMonitor.RequiredConfirmationBlocks)

		self.blockDownloader = listener.NewBlockDownloader(self.config).
			WithClient(client).
			WithInputChannel(networkMonitor.Output).
			WithMonitor(monitor).
			WithBackoff(0, self.config.Syncer.TransactionMaxInterval)

		if self.lastAcceptedArweaveHeight > 0 {
			// This is a restart from the watchdog so set the start height
			// Otherwise it will be set later
			self.blockDownloader.WithHeightRange(self.lastAcceptedArweaveHeight+1, math.MaxUint64)
		}

		transactionDownloader := listener.NewTransactionDownloader(self.config).
			WithClient(client).
			WithInputChannel(self.blockDownloader.Output).
			WithMonitor(monitor).
			WithBackoff(0, self.config.Syncer.TransactionMaxInterval).
			WithFilterInteractions()

		store := NewStore(self.config).
			WithInputChannel(transactionDownloader.Output).
			WithMonitor(monitor)
		self.store = store

		return task.NewTask(self.config, "watched").
			WithSubtask(networkMonitor.Task).
			WithSubtask(self.blockDownloader.Task).
			WithSubtask(transactionDownloader.Task).
			WithSubtask(store.Task)
	}

	watchdog := task.NewWatchdog(self.config).
		WithTask(watched).
		WithIsOK(30*time.Second, func() bool {
			isOK := monitor.IsOK()
			if !isOK {
				monitor.Clear()
				monitor.GetReport().Run.Errors.NumWatchdogRestarts.Inc()
			}
			return isOK
		})

	self.Task = self.Task.
		WithSubtask(server.Task).
		WithSubtask(monitor.Task).
		WithConditionalSubtask(self.config.Syncer.Enabled, watchdog.Task)

	// Starts all the tasks, but downloading new block will be blocked until lastAcceptedArweaveHeight is set
	err = self.Start()
	if err != nil {
		return
	}

	return
}

func (self *SyncerController) isRunning() bool {
	return self != nil && self.config.Syncer.Enabled && self.Task != nil && !self.Task.IsStopping.Load()
}

func (self *SyncerController) GetNextArweaveBlock(height uint64) *types.NextArweaveBlock {
	if !self.isRunning() {
		return nil
	}
	return self.store.GetNextArweaveBlock(height)
}

func (self *SyncerController) StopWait() {
	if self == nil {
		return
	}

	self.StopWait()
}

func (self *SyncerController) SetLastAcceptedBlockHeight(height uint64) {
	if !self.isRunning() {
		return
	}
	self.mtx.Lock()
	defer self.mtx.Unlock()

	if self.lastAcceptedArweaveHeight == 0 {
		// This is the first time we set the height
		self.blockDownloader.SetStartHeight(self.lastAcceptedArweaveHeight + 1)
	} else {
		// Called after the initialization
		self.store.RemoveNextArweaveBlocksUpToHeight(height)

	}
}
