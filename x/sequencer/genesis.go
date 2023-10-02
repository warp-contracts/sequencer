package sequencer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/config"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, configProvider *config.ConfigProvider) {
	// Set LastArweaveBlock
	var lastArweaveBlock *types.LastArweaveBlock
	var err error
	if genState.LastArweaveBlock != nil {
		lastArweaveBlock = genState.LastArweaveBlock
	} else {
		lastArweaveBlock, err = configProvider.LastArweaveBlock()
		if err != nil {
			panic(err)
		}
	}
	k.SetLastArweaveBlock(ctx, *lastArweaveBlock)

	// Set all the lastSortKey
	var lastSortKeys []types.LastSortKey
	if len(genState.LastSortKeyList) == 0 {
		lastSortKeys = configProvider.LastSortKeys()
	} else {
		lastSortKeys = genState.GetLastSortKeyList()
	}
	for _, elem := range lastSortKeys {
		k.SetLastSortKey(ctx, elem)
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
	genesis.LastSortKeyList = k.GetAllLastSortKey(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
