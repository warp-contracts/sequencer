package sequencer

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

const (
	ARWEAVE_BLOCK_FILE  = "arweave_block.json"
	PREV_SORT_KEYS_FILE = "prev_sort_keys.json"
)

// Loads the genesis state from configuration files
type GenesisLoader interface {
	LoadArweaveBlock() *types.GenesisArweaveBlock
	LoadPrevSortKeys() []types.PrevSortKey
}

type GenesisConfigFileLoader struct {
	logger      log.Logger
	genesisPath string
}

var homePath string

func NewGenesisLoader(logger log.Logger) GenesisLoader {
	genesisPath := path.Join(homePath, "genesis")
	return &GenesisConfigFileLoader{
		logger,
		genesisPath,
	}
}

func ProvideGenesisLoader(defaultHomeNode string) func(logger log.Logger) GenesisLoader {
	homePath = defaultHomeNode
	return NewGenesisLoader
}

func (loader *GenesisConfigFileLoader) LoadArweaveBlock() *types.GenesisArweaveBlock {
	filePath := filepath.Join(loader.genesisPath, ARWEAVE_BLOCK_FILE)
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		loader.
			logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to retrieve arweave block from the file")
		return nil
	}

	var block types.GenesisArweaveBlock
	err = json.Unmarshal(jsonFile, &block)
	if err != nil {
		loader.
			logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to unmarshal arweave block from the file")
		return nil
	}

	loader.
		logger.
		With("file", filePath).
		With("last arweave block height", block.LastArweaveBlock.Height).
		With("next arweave block height", block.NextArweaveBlock.BlockInfo.Height).
		Info("Arweave block loaded from the file")

	return &block
}

func (loader *GenesisConfigFileLoader) LoadPrevSortKeys() []types.PrevSortKey {
	filePath := filepath.Join(loader.genesisPath, PREV_SORT_KEYS_FILE)
	var keys []types.PrevSortKey

	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		loader.
			logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to retrieve prev sort keys from the file")
		return keys
	}

	err = json.Unmarshal(jsonFile, &keys)
	if err != nil {
		loader.
			logger.
			With("err", err).
			With("file", filePath).
			Info("Unable to unmarshal prev sort keys from the file")
		return keys
	}

	loader.
		logger.
		With("number of keys", len(keys)).
		With("file", filePath).
		Info("Prev sort keys loaded from the file")
	return keys
}

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, loader GenesisLoader) {
	// Set LastArweaveBlock
	var lastArweaveBlock *types.LastArweaveBlock
	if genState.LastArweaveBlock != nil {
		lastArweaveBlock = genState.LastArweaveBlock
	} else {
		block := loader.LoadArweaveBlock()
		if block == nil {
			panic("A sequencer blockchain cannot launch without a genesis Arweave block")
		}
		lastArweaveBlock = &types.LastArweaveBlock{
			ArweaveBlock:         block.LastArweaveBlock,
			SequencerBlockHeight: 0,
		}
	}
	k.SetLastArweaveBlock(ctx, *lastArweaveBlock)

	// Set all the prevSortKey
	var prevSortKeys []types.PrevSortKey
	if len(genState.PrevSortKeyList) == 0 {
		prevSortKeys = loader.LoadPrevSortKeys()
	} else {
		prevSortKeys = genState.GetPrevSortKeyList()
	}
	for _, elem := range prevSortKeys {
		k.SetPrevSortKey(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	err := k.SetParams(ctx, genState.Params)
	if err != nil {
		panic(err)
	}
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
	genesis.PrevSortKeyList = k.GetAllPrevSortKey(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
