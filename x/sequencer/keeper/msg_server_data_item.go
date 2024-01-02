package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k *msgServer) DataItem(goCtx context.Context, msg *types.MsgDataItem) (*types.MsgDataItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Find the contract id from the tags
	contract, err := msg.GetContractFromTags()
	if err != nil {
		return nil, err
	}

	// Assign prevSortKey
	prevSortKey := types.PrevSortKey{
		Contract: contract,
		SortKey:  msg.SortKey,
	}

	k.SetPrevSortKey(ctx, prevSortKey)

	// Increment the limiter counter that corresponds to the contract
	k.limiterKeeper.Inc(ctx, 0, []byte(contract))

	// L2 interactions contained in the last sequencer block
	k.blockInteractions.Add(msg)

	return &types.MsgDataItemResponse{}, nil
}
