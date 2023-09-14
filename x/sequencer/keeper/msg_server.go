package keeper

import (
	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type msgServer struct {
	Keeper

	Controller controller.ArweaveBlocksController
	lastSortKey *types.SortKey
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, controller controller.ArweaveBlocksController) types.MsgServer {
	return &msgServer{Keeper: keeper, Controller: controller}
}

var _ types.MsgServer = &msgServer{}
