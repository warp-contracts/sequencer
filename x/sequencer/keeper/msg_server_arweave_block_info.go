package keeper

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) ArweaveBlockInfo(goCtx context.Context, msg *types.MsgArweaveBlockInfo) (*types.MsgArweaveBlockInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var newValue = &types.ArweaveBlockInfo{
		Height:    msg.Height,
		Timestamp: msg.Timestamp,
		Hash:      msg.Hash,
	}

	if err := k.checkBlockIsOldEnough(ctx, newValue); err != nil {
		return nil, err
	}

	if err := k.compareBlockWithPreviousOne(ctx, newValue); err != nil {
		return nil, err
	}

	if err := k.compareWithNextBlockAndRemove(ctx, newValue); err != nil {
		return nil, err
	}

	k.SetLastArweaveBlock(ctx, *newValue)
	return &types.MsgArweaveBlockInfoResponse{}, nil
}

func (k msgServer) checkBlockIsOldEnough(ctx sdk.Context, newValue *types.ArweaveBlockInfo) error {
	arweaveBlockTimestamp := time.Unix(int64(newValue.Timestamp), 0)
	cosmosBlockTimestamp := ctx.BlockHeader().Time

	if cosmosBlockTimestamp.Before(arweaveBlockTimestamp.Add(time.Hour)) {
		return errors.Wrapf(types.ErrArweaveBlockTimestampMismatch,
			"The timestamp of the Arweave block (%s) should be one hour earlier than the Cosmos block (%s)",
			arweaveBlockTimestamp.UTC(), cosmosBlockTimestamp.UTC())
	}
	return nil
}

func (k msgServer) compareBlockWithPreviousOne(ctx sdk.Context, newValue *types.ArweaveBlockInfo) error {
	oldValue, isFound := k.GetLastArweaveBlock(ctx)

	if isFound {
		if newValue.Height-oldValue.Height != 1 {
			return errors.Wrap(types.ErrArweaveBlockHeightMismatch,
				"The new height of the Arweave block is not the next value compared to the previous height")
		}

		if newValue.Timestamp <= oldValue.Timestamp {
			return errors.Wrap(types.ErrArweaveBlockTimestampMismatch,
				"The timestamp of the Arweave block is not later than the previous one")
		}
	}
	return nil
}

func (k msgServer) compareWithNextBlockAndRemove(ctx sdk.Context, newValue *types.ArweaveBlockInfo) error {
	heightStr := strconv.FormatUint(newValue.Height, 10)
	nextArweaveBlock, isFound := k.GetNextArweaveBlock(ctx, heightStr)
	if isFound {
		if newValue.Timestamp != nextArweaveBlock.BlockInfo.Timestamp {
			return errors.Wrap(types.ErrArweaveBlockTimestampMismatch,
				"The timestamp of the Arweave block does not match the timestamp of the block downloaded by the Validator")
		}
		if !bytes.Equal(newValue.Hash, nextArweaveBlock.BlockInfo.Hash) {
			return errors.Wrap(types.ErrArweaveBlockHashMismatch,
				"The hash of the Arweave block does not match the hash of the block downloaded by the Validator")
		}

		k.RemoveNextArweaveBlock(ctx, heightStr)
		return nil
	} else {
		return errors.Wrapf(types.ErrInvalidArweaveBlockTx,
			"The validator did not fetch the Arweave block at height %s", heightStr)
	}
}
