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

		LastArweaveBlock: &types.LastArweaveBlock{},
		PrevSortKeyList: []types.PrevSortKey{
			{
				Contract: "0",
			},
			{
				Contract: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SequencerKeeper(t)
	sequencer.InitGenesis(ctx, *k, genesisState, nil)
	got := sequencer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.LastArweaveBlock, got.LastArweaveBlock)
	require.ElementsMatch(t, genesisState.PrevSortKeyList, got.PrevSortKeyList)
	// this line is used by starport scaffolding # genesis/test/assert
}
