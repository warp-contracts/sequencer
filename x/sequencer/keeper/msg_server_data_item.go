package keeper

import (
	"context"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) Arweave(goCtx context.Context, msg *types.MsgArweave) (*types.MsgArweaveResponse, error) {
	return &types.MsgArweaveResponse{}, nil
}
