package ante

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// It checks if the transaction representing an Arweave block contains exacly one message: MsgArweaveBlock
// Additionally, it checks if the Cosmos block does not lack transactions with the Arweave block.
type ArweaveBlockTxDecorator struct {
	keeper     keeper.Keeper
	controller controller.ArweaveBlocksController
}

func NewArweaveBlockTxDecorator(keeper keeper.Keeper, controller controller.ArweaveBlocksController) ArweaveBlockTxDecorator {
	return ArweaveBlockTxDecorator{keeper, controller}
}

func (abtd ArweaveBlockTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	foundTx, err := verifyArweaveBlockTx(tx)
	if err != nil {
		return ctx, err
	}

	if foundTx {
		if ctx.IsCheckTx() {
			return ctx, errors.Wrapf(types.ErrArweaveBlockNotFromProposer,
				"transaction with arweave block can only be added by the Proposer")
		}

		// Valid transaction with Arweave block does not require further validation within AnteHandler
		return ctx, nil
	} else if err := abtd.shouldBlockContainArweaveTx(ctx); err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func verifyArweaveBlockTx(tx sdk.Tx) (bool, error) {
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		arweaveBlock, isArweaveBlock := msg.(*types.MsgArweaveBlock)
		if isArweaveBlock {
			if len(msgs) > 1 {
				return true, errors.Wrapf(types.ErrTooManyMessages,
					"transaction with arweave block can have only one message, and it has: %d", len(msgs))
			}
			return true, arweaveBlock.ValidateBasic()
		}
	}
	return false, nil
}

func (abtd ArweaveBlockTxDecorator) shouldBlockContainArweaveTx(ctx sdk.Context) error {
	if ctx.BlockHeader().Height == 0 || ctx.IsCheckTx() {
		return nil
	}

	lastArweaveBlock := abtd.keeper.MustGetLastArweaveBlock(ctx)
	nextArweaveBlock := abtd.controller.GetNextArweaveBlock(lastArweaveBlock.Height)

	if nextArweaveBlock != nil && types.IsArweaveBlockOldEnough(ctx, nextArweaveBlock.BlockInfo) {
		return errors.Wrapf(types.ErrNoArweaveBlockTx,
			"The first transaction of the block should contain a transaction with the Arweave block with height %d",
			lastArweaveBlock.Height)
	}
	return nil
}

func isArweaveBlockTx(tx sdk.Tx) bool {
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		_, isArweaveBlock := msg.(*types.MsgArweaveBlock)
		if isArweaveBlock {
			return true
		}
	}
	return false
}
