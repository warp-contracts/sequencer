package ante

import (
	"fmt"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// It checks if the transaction representing an Arweave block contains exacly one message: MsgArweaveBlock
type ArweaveBlockTxDecorator struct {
	keeper keeper.Keeper
}

func NewArweaveBlockTxDecorator(keeper keeper.Keeper) ArweaveBlockTxDecorator {
	return ArweaveBlockTxDecorator{keeper}
}

func (abtd ArweaveBlockTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	foundTx, err := verifyArweaveBlockTx(tx)
	if err != nil {
		return ctx, err
	}

	if ctx.IsReCheckTx() && !foundTx {
		if err := shouldBlockContainArweaveTx(ctx, &abtd.keeper); err != nil {
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

func shouldBlockContainArweaveTx(ctx sdk.Context, keeper *keeper.Keeper) error {
	lastArweaveBlock, _ := keeper.GetLastArweaveBlock(ctx)
	nextHeight := fmt.Sprintf("%d", lastArweaveBlock.Height+1)
	nextArweaveBlock, found := keeper.GetNextArweaveBlock(ctx, nextHeight)

	if found {
		nextArweaveBlockTimestamp := time.Unix(int64(nextArweaveBlock.BlockInfo.Timestamp), 0)
		cosmosBlockTimestamp := ctx.BlockHeader().Time
		if cosmosBlockTimestamp.After(nextArweaveBlockTimestamp.Add(time.Hour)) {
			return errors.Wrapf(types.ErrNoArweaveBlockTx,
				"The first transaction of the block should contain a transaction with the Arweave block with height %s",
				nextHeight)
		}
	}
	return nil
}
