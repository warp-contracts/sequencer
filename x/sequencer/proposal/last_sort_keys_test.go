package proposal

import (
	"github.com/stretchr/testify/require"
	"testing"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestGetLastSortKeyNoKey(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	lastSortKeys := newLastSortKeys(k, ctx)

	key := lastSortKeys.getLastSortKey("abc")

	require.Equal(t, "", key)
}

func TestGetLastSortKeyKeyInStore(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	k.SetLastSortKey(ctx, types.LastSortKey{
		Contract: "abc",
		SortKey: "123",
	})
	lastSortKeys := newLastSortKeys(k, ctx)

	key := lastSortKeys.getLastSortKey("abc")

	require.Equal(t, "123", key)
}

func TestGetLastSortKeyKeyInMemory(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	lastSortKeys := newLastSortKeys(k, ctx)

	key := lastSortKeys.getAndStoreLastSortKey("abc", "123")
	require.Equal(t, "", key)

	key = lastSortKeys.getLastSortKey("abc")
	require.Equal(t, "123", key)
}

func TestGetLastSortKeyKeyInMemoryAndStore(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	k.SetLastSortKey(ctx, types.LastSortKey{
		Contract: "abc",
		SortKey: "123",
	})
	lastSortKeys := newLastSortKeys(k, ctx)

	key := lastSortKeys.getAndStoreLastSortKey("abc", "456")
	require.Equal(t, "123", key)

	key = lastSortKeys.getLastSortKey("abc")
	require.Equal(t, "456", key)
}