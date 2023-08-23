package keeper

import (
	"context"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) ArweaveTransaction(goCtx context.Context, msg *types.MsgArweaveTransaction) (*types.MsgArweaveTransactionResponse, error) {
	return &types.MsgArweaveTransactionResponse{}, nil
}
