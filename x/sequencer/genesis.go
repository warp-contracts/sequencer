package sequencer

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/cometbft/cometbft/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

const (
	LAST_ARWEAVE_BLOCK_FILE = "last_arweave_block.json"
	LAST_SORT_KEYS_FILE     = "last_sort_keys.json"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, configPath string) {
	// Set LastArweaveBlock
	var lastArweaveBlock *types.LastArweaveBlock
	var err error
	if genState.LastArweaveBlock != nil {
		lastArweaveBlock = genState.LastArweaveBlock
	} else {
		lastArweaveBlock, err = readLastArweaveBlockFromFile(ctx.Logger(), configPath)
		if err != nil {
			panic(err)
		}
	}
	k.SetLastArweaveBlock(ctx, *lastArweaveBlock)

	// Set all the lastSortKey
	var lastSortKeys []types.LastSortKey
	if len(genState.LastSortKeyList) == 0 {
		lastSortKeys = readLastSortKeysFromFile(ctx.Logger(), configPath)
	} else {
		lastSortKeys = genState.GetLastSortKeyList()
	}
	for _, elem := range lastSortKeys {
		k.SetLastSortKey(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

func readLastArweaveBlockFromFile(logger log.Logger, configPath string) (*types.LastArweaveBlock, error) {
	filePath := filepath.Join(configPath, LAST_ARWEAVE_BLOCK_FILE)
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var blockInfo types.ArweaveBlockInfo
	err = json.Unmarshal(jsonFile, &blockInfo)
	if err != nil {
		return nil, err
	}

	return &types.LastArweaveBlock{
		ArweaveBlock:         &blockInfo,
		SequencerBlockHeight: 0,
	}, nil
}

func readLastSortKeysFromFile(logger log.Logger, configPath string) []types.LastSortKey {
	filePath := filepath.Join(configPath, LAST_SORT_KEYS_FILE)
	var keys []types.LastSortKey

	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to retrieve last sort keys from the file")
		return keys
	}

	err = json.Unmarshal(jsonFile, &keys)
	if err != nil {
		logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to unmarshal last sort keys from the file")
		return keys
	}

	logger.
		With("number of keys", len(keys)).
		With("file", filePath).
		Info("Last sort keys loaded from the file")
	return keys
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
