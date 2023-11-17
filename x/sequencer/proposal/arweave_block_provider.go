package proposal

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer"
	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Returns the last Arweave block already added to the blockchain and the next Arweave block that can be added
// In the case of the first sequencer block, it loads an Arweave block from the genesis state
type ArweaveBlockProvider struct {
	keeper     *keeper.Keeper
	controller controller.ArweaveBlocksController
	loader     sequencer.GenesisLoader
}

func NewArweaveBlockProvider(keeper *keeper.Keeper, controller controller.ArweaveBlocksController, loader sequencer.GenesisLoader) *ArweaveBlockProvider {
	return &ArweaveBlockProvider{
		keeper,
		controller,
		loader,
	}
}

func (provider *ArweaveBlockProvider) getLastArweaveBlock(ctx sdk.Context, firstBlock bool) types.LastArweaveBlock {
	if firstBlock {
		genesisBlock := provider.loader.LoadArweaveBlock()
		if genesisBlock != nil {
			return types.LastArweaveBlock{
				ArweaveBlock:         genesisBlock.LastArweaveBlock,
				SequencerBlockHeight: 0,
			}
		}
	}
	return provider.keeper.MustGetLastArweaveBlock(ctx)
}

func (provider *ArweaveBlockProvider) getNextArweaveBlock(height uint64, firstBlock bool) *types.NextArweaveBlock {
	if firstBlock {
		genesisBlock := provider.loader.LoadArweaveBlock()
		if genesisBlock != nil {
			return genesisBlock.NextArweaveBlock
		}
	}
	return provider.controller.GetNextArweaveBlock(height)
}
