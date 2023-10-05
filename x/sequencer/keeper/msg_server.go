package keeper

import (
	"github.com/warp-contracts/sequencer/x/sequencer/ante"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type msgServer struct {
	Keeper
	blockInteractions *ante.BlockInteractions
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, blockInteractions *ante.BlockInteractions) types.MsgServer {
	return &msgServer{keeper, blockInteractions}
}

var _ types.MsgServer = &msgServer{}
