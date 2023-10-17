package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// SetPrevSortKey set a specific prevSortKey in the store from its index
func (k Keeper) SetPrevSortKey(ctx sdk.Context, prevSortKey types.PrevSortKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrevSortKeyKeyPrefix))
	b := k.cdc.MustMarshal(&prevSortKey)
	store.Set(types.PrevSortKeyKey(
		prevSortKey.Contract,
	), b)
}

// GetPrevSortKey returns a prevSortKey from its index
func (k Keeper) GetPrevSortKey(
	ctx sdk.Context,
	contract string,

) (val types.PrevSortKey, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrevSortKeyKeyPrefix))

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
	ctx sdk.Context,
	contract string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrevSortKeyKeyPrefix))
	store.Delete(types.PrevSortKeyKey(
		contract,
	))
}

// GetAllPrevSortKey returns all prevSortKey
func (k Keeper) GetAllPrevSortKey(ctx sdk.Context) (list []types.PrevSortKey) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrevSortKeyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PrevSortKey
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
