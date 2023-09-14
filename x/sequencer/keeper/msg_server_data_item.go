package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k *msgServer) DataItem(goCtx context.Context, msg *types.MsgDataItem) (*types.MsgDataItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.validateSortKey(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgDataItemResponse{}, nil
}

func (k *msgServer) validateSortKey(ctx sdk.Context, msg *types.MsgDataItem) error {
	if k.lastSortKey == nil || k.lastSortKey.SequencerHeight != ctx.BlockHeight() {
		k.lastSortKey = types.NewSortKey(k.MustGetLastArweaveBlock(ctx).ArweaveBlock.Height, ctx.BlockHeight())
	}

	expectedSortKey := k.lastSortKey.GetNextValue()
	if expectedSortKey != msg.SortKey {
		return errors.Wrapf(types.ErrInvalidSortKey,
			"Invalid sort key. Expected: %s, received: %s", expectedSortKey, msg.SortKey)
	}
	return nil
}
