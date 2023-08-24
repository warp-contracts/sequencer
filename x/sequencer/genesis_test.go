package sequencer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/testutil/nullify"
	"github.com/warp-contracts/sequencer/x/sequencer"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LastArweaveBlock: &types.ArweaveBlockInfo{},
		NextArweaveBlockList: []types.NextArweaveBlock{
			{
				BlockInfo: &types.ArweaveBlockInfo{
					Height: 0,
				},
			},
			{
				BlockInfo: &types.ArweaveBlockInfo{
					Height: 1,
				},
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SequencerKeeper(t)
	sequencer.InitGenesis(ctx, *k, genesisState)
	got := sequencer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.LastArweaveBlock, got.LastArweaveBlock)
	require.ElementsMatch(t, genesisState.NextArweaveBlockList, got.NextArweaveBlockList)
	// this line is used by starport scaffolding # genesis/test/assert
}
