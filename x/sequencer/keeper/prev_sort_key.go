package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// SetPrevSortKey set a specific prevSortKey in the store from its index
func (k Keeper) SetPrevSortKey(ctx context.Context, prevSortKey types.PrevSortKey) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PrevSortKeyKeyPrefix))
	b := k.cdc.MustMarshal(&prevSortKey)
	store.Set(types.PrevSortKeyKey(
		prevSortKey.Contract,
	), b)
}

// GetPrevSortKey returns a prevSortKey from its index
func (k Keeper) GetPrevSortKey(
	ctx context.Context,
	contract string,

) (val types.PrevSortKey, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PrevSortKeyKeyPrefix))

	b := store.Get(types.PrevSortKeyKey(
		contract,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePrevSortKey removes a prevSortKey from the store
func (k Keeper) RemovePrevSortKey(
	ctx context.Context,
	contract string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PrevSortKeyKeyPrefix))
	store.Delete(types.PrevSortKeyKey(
		contract,
	))
}

// GetAllPrevSortKey returns all prevSortKey
func (k Keeper) GetAllPrevSortKey(ctx context.Context) (list []types.PrevSortKey) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PrevSortKeyKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PrevSortKey
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
