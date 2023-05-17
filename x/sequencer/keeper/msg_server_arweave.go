package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) Arweave(goCtx context.Context, msg *types.MsgArweave) (*types.MsgArweaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgArweaveResponse{}, nil
}
