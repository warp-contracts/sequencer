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

	if ctx.IsReCheckTx() && !foundTx {
		if err := abtd.shouldBlockContainArweaveTx(ctx); err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}

func verifyArweaveBlockTx(tx sdk.Tx) (bool, error) {
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		_, isArweaveBlock := msg.(*types.MsgArweaveBlock)
		if isArweaveBlock {
			if len(msgs) > 1 {
				return true, errors.Wrapf(types.ErrTooManyMessages,
					"transaction with arweave block can have only one message, and it has: %d", len(msgs))
			}
			return true, nil
		}
	}
	return false, nil
}

func (abtd ArweaveBlockTxDecorator) shouldBlockContainArweaveTx(ctx sdk.Context) error {
	lastArweaveBlock := abtd.keeper.MustGetLastArweaveBlock(ctx)
	nextArweaveBlock := abtd.controller.GetNextArweaveBlock(lastArweaveBlock.Height)

	if nextArweaveBlock != nil && types.CheckArweaveBlockIsOldEnough(ctx, nextArweaveBlock.BlockInfo) {
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
