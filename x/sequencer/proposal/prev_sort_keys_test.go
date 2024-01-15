package proposal

import (
	"github.com/stretchr/testify/require"
	"testing"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestGetPrevSortKeyNoKey(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	prevSortKeys := newPrevSortKeys(&k, ctx)

	key := prevSortKeys.getPrevSortKey("abc")

	require.Equal(t, "", key)
}

func TestGetPrevSortKeyKeyInStore(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	k.SetPrevSortKey(ctx, types.PrevSortKey{
		Contract: "abc",
		SortKey: "123",
	})
	prevSortKeys := newPrevSortKeys(&k, ctx)

	key := prevSortKeys.getPrevSortKey("abc")

	require.Equal(t, "123", key)
}

func TestGetPrevSortKeyKeyInMemory(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	prevSortKeys := newPrevSortKeys(&k, ctx)

	key := prevSortKeys.getAndStorePrevSortKey("abc", "123")
	require.Equal(t, "", key)

	key = prevSortKeys.getPrevSortKey("abc")
	require.Equal(t, "123", key)
}

func TestGetPrevSortKeyKeyInMemoryAndStore(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	k.SetPrevSortKey(ctx, types.PrevSortKey{
		Contract: "abc",
		SortKey: "123",
	})
	prevSortKeys := newPrevSortKeys(&k, ctx)

	key := prevSortKeys.getAndStorePrevSortKey("abc", "456")
	require.Equal(t, "123", key)

	key = prevSortKeys.getPrevSortKey("abc")
	require.Equal(t, "456", key)
}