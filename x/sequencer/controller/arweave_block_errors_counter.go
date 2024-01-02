package controller

import (
	"github.com/cometbft/cometbft/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
)

// Stores the number of consecutive errors related to the Arweave blocks
// and in the case of a long series, restarts the controller.
type ArweaveBlockErrorsCounter struct {
	controller                    ArweaveBlocksController
	keeper                        *keeper.Keeper
	logger                        log.Logger
	consecutiveArweaveBlockErrors int
}

func NewArweaveBlockErrorsCounter(controller ArweaveBlocksController, keeper *keeper.Keeper, logger log.Logger) *ArweaveBlockErrorsCounter {
	return &ArweaveBlockErrorsCounter{
		controller:                    controller,
		keeper:                        keeper,
		logger:                        logger.With("module", "block-validator"),
		consecutiveArweaveBlockErrors: 0,
	}
}

func (errors *ArweaveBlockErrorsCounter) Reset() {
	if errors.consecutiveArweaveBlockErrors > 0 {
		errors.logger.With("consecutive_errors", errors.consecutiveArweaveBlockErrors).
			Debug("End of the series of Arweave block errors")
	}
	errors.consecutiveArweaveBlockErrors = 0
}

func (errors *ArweaveBlockErrorsCounter) Error(ctx sdk.Context) {
	errors.consecutiveArweaveBlockErrors++
	errors.logger.
		With("consecutive_errors", errors.consecutiveArweaveBlockErrors).
		Debug("Invalid Arweave block error")

	// setting the block in case the controller is not initialized
	errors.controller.SetLastAcceptedBlock(errors.keeper.MustGetLastArweaveBlock(ctx).ArweaveBlock)

	if errors.consecutiveArweaveBlockErrors > 10 {
		errors.consecutiveArweaveBlockErrors = 0
		errors.logger.Error("Controller restart due to too many consecutive Arweave block errors")
		errors.controller.Restart()
	}
}
