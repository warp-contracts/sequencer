package keeper

import (
	"context"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) LastArweaveBlock(goCtx context.Context, msg *types.MsgLastArweaveBlock) (*types.MsgLastArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var newValue = types.LastArweaveBlock{
		Creator:   msg.Creator,
		Height:    msg.Height,
		Timestamp: msg.Timestamp,
		Hash:      msg.Hash,
	}

	arweaveBlockTimestamp := time.Unix(int64(newValue.Timestamp), 0)
	cosmosBlockTimestamp := ctx.BlockHeader().Time

	if cosmosBlockTimestamp.Before(arweaveBlockTimestamp.Add(time.Hour)) {
		return nil, errors.Wrapf(types.ErrArweaveBlockTimestampMismatch,
			"The timestamp of the Arweave block (%s) should be one hour earlier than the Cosmos block (%s)",
			arweaveBlockTimestamp.UTC(), cosmosBlockTimestamp.UTC())
	}

	oldValue, isFound := k.GetLastArweaveBlock(ctx)

	if isFound {
		if newValue.Height-oldValue.Height != 1 {
			return nil, errors.Wrap(types.ErrArweaveBlockHeightMismatch,
				"The new height of the Arweave block is not the next value compared to the previous height")
		}

		if newValue.Timestamp <= oldValue.Timestamp {
			return nil, errors.Wrap(types.ErrArweaveBlockTimestampMismatch,
				"The timestamp of the Arweave block is not later than the previous one")
		}
	}

	k.SetLastArweaveBlock(
		ctx,
		newValue)
	return &types.MsgLastArweaveBlockResponse{}, nil
}
