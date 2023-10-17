package proposal

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
)

// Allows calculating the value of PrevSortKey. 
// It get the current values from the store, and for calculating/validation, it keeps new values in memory.
type PrevSortKeys struct {
	keeper       *keeper.Keeper
	ctx          sdk.Context
	prevSortKeys map[string]string
}

func newPrevSortKeys(keeper *keeper.Keeper, ctx sdk.Context) *PrevSortKeys {
	return &PrevSortKeys{
		keeper:       keeper,
		ctx:          ctx,
		prevSortKeys: make(map[string]string),
	}
}

func (keys *PrevSortKeys) getPrevSortKey(contract string) string {
	key := keys.prevSortKeys[contract]
	if key != "" {
		return key
	}
	prevSortKey, found := keys.keeper.GetPrevSortKey(keys.ctx, contract)
	if found {
		return prevSortKey.SortKey
	}
	return ""
}

func (keys *PrevSortKeys) getAndStorePrevSortKey(contract string, sortKey string) string {
	key := keys.getPrevSortKey(contract)
	keys.prevSortKeys[contract] = sortKey
	return key
}
