package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/testutil/nullify"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func createTestLastArweaveBlock(keeper *keeper.Keeper, ctx sdk.Context) types.LastArweaveBlock {
	item := types.LastArweaveBlock{}
	keeper.SetLastArweaveBlock(ctx, item)
	return item
}

func TestLastArweaveBlockGet(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	item := createTestLastArweaveBlock(keeper, ctx)
	rst, found := keeper.GetLastArweaveBlock(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLastArweaveBlockRemove(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	createTestLastArweaveBlock(keeper, ctx)
	keeper.RemoveLastArweaveBlock(ctx)
	_, found := keeper.GetLastArweaveBlock(ctx)
	require.False(t, found)
}
