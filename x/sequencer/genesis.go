package sequencer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.LastArweaveBlock != nil {
		k.SetLastArweaveBlock(ctx, *genState.LastArweaveBlock)
	}
	// Set all the nextArweaveBlock
	for _, elem := range genState.NextArweaveBlockList {
		k.SetNextArweaveBlock(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all lastArweaveBlock
	lastArweaveBlock, found := k.GetLastArweaveBlock(ctx)
	if found {
		genesis.LastArweaveBlock = &lastArweaveBlock
	}
	genesis.NextArweaveBlockList = k.GetAllNextArweaveBlock(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
