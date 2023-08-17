package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) LastArweaveBlock(goCtx context.Context, msg *types.MsgLastArweaveBlock) (*types.MsgLastArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var lastArweaveBlock = types.LastArweaveBlock{
		Creator: msg.Creator,
	}

	k.SetLastArweaveBlock(
		ctx,
		lastArweaveBlock,
	)
	return &types.MsgLastArweaveBlockResponse{}, nil
}
