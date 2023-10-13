package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k *msgServer) DataItem(goCtx context.Context, msg *types.MsgDataItem) (*types.MsgDataItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.setContractLastSortKey(ctx, msg); err != nil {
		return nil, err
	}
	
	k.blockInteractions.Add(msg)

	return &types.MsgDataItemResponse{}, nil
}

func (k *msgServer) setContractLastSortKey(ctx sdk.Context, msg *types.MsgDataItem) error {
	contract, err := msg.GetContractFromTags()
	if err != nil {
		return err
	}

	lastSortKey := types.LastSortKey{
		Contract: contract,
		SortKey:  msg.SortKey,
	}

	k.SetLastSortKey(ctx, lastSortKey)
	return nil
}
