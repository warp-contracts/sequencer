package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// SetLastSortKey set a specific lastSortKey in the store from its index
func (k Keeper) SetLastSortKey(ctx sdk.Context, lastSortKey types.LastSortKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastSortKeyKeyPrefix))
	b := k.cdc.MustMarshal(&lastSortKey)
	store.Set(types.LastSortKeyKey(
		lastSortKey.Contract,
	), b)
}

// GetLastSortKey returns a lastSortKey from its index
func (k Keeper) GetLastSortKey(
	ctx sdk.Context,
	contract string,

) (val types.LastSortKey, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastSortKeyKeyPrefix))

	b := store.Get(types.LastSortKeyKey(
		contract,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLastSortKey removes a lastSortKey from the store
func (k Keeper) RemoveLastSortKey(
	ctx sdk.Context,
	contract string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastSortKeyKeyPrefix))
	store.Delete(types.LastSortKeyKey(
		contract,
	))
}

// GetAllLastSortKey returns all lastSortKey
func (k Keeper) GetAllLastSortKey(ctx sdk.Context) (list []types.LastSortKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastSortKeyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LastSortKey
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
