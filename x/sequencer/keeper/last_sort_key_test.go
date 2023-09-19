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

func createNLastSortKey(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LastSortKey {
	items := make([]types.LastSortKey, n)
	for i := range items {
		items[i].Contract = strconv.Itoa(i)

		keeper.SetLastSortKey(ctx, items[i])
	}
	return items
}

func TestLastSortKeyGet(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNLastSortKey(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLastSortKey(ctx,
			item.Contract,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLastSortKeyRemove(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNLastSortKey(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLastSortKey(ctx,
			item.Contract,
		)
		_, found := keeper.GetLastSortKey(ctx,
			item.Contract,
		)
		require.False(t, found)
	}
}

func TestLastSortKeyGetAll(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNLastSortKey(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLastSortKey(ctx)),
	)
}
