package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/testutil/nullify"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPrevSortKey(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PrevSortKey {
	items := make([]types.PrevSortKey, n)
	for i := range items {
		items[i].Contract = strconv.Itoa(i)

		keeper.SetPrevSortKey(ctx, items[i])
	}
	return items
}

func TestPrevSortKeyGet(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNPrevSortKey(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPrevSortKey(ctx,
			item.Contract,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPrevSortKeyRemove(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNPrevSortKey(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePrevSortKey(ctx,
			item.Contract,
		)
		_, found := keeper.GetPrevSortKey(ctx,
			item.Contract,
		)
		require.False(t, found)
	}
}

func TestPrevSortKeyGetAll(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNPrevSortKey(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPrevSortKey(ctx)),
	)
}
