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

func createNNextArweaveBlock(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NextArweaveBlock {
	items := make([]types.NextArweaveBlock, n)
	for i := range items {
		items[i].BlockInfo = &types.ArweaveBlockInfo{
			Height: uint64(i),
		}

		keeper.SetNextArweaveBlock(ctx, items[i])
	}
	return items
}

func TestNextArweaveBlockGet(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNNextArweaveBlock(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNextArweaveBlock(ctx,
			item.GetHeightString(),
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNextArweaveBlockRemove(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNNextArweaveBlock(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNextArweaveBlock(ctx,
			item.GetHeightString(),
		)
		_, found := keeper.GetNextArweaveBlock(ctx,
			item.GetHeightString(),
		)
		require.False(t, found)
	}
}

func TestNextArweaveBlockGetAll(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNNextArweaveBlock(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNextArweaveBlock(ctx)),
	)
}
