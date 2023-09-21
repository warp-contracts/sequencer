package proposal

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
)

// Allows calculating the value of LastSortKey. 
// It get the current values from the store, and for calculating/validation, it keeps new values in memory.
type LastSortKeys struct {
	keeper       *keeper.Keeper
	ctx          sdk.Context
	lastSortKeys map[string]string
}

func newLastSortKeys(keeper *keeper.Keeper, ctx sdk.Context) *LastSortKeys {
	return &LastSortKeys{
		keeper:       keeper,
		ctx:          ctx,
		lastSortKeys: make(map[string]string),
	}
}

func (keys *LastSortKeys) getLastSortKey(contract string) string {
	key := keys.lastSortKeys[contract]
	if key != "" {
		return key
	}
	lastSortKey, found := keys.keeper.GetLastSortKey(keys.ctx, contract)
	if found {
		return lastSortKey.SortKey
	}
	return ""
}

func (keys *LastSortKeys) getAndStoreLastSortKey(contract string, sortKey string) string {
	key := keys.getLastSortKey(contract)
	keys.lastSortKeys[contract] = sortKey
	return key
}
