package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) DataItem(goCtx context.Context, msg *types.MsgDataItem) (*types.MsgDataItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDataItemResponse{}, nil
}
