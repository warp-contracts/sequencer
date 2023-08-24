package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// SetNextArweaveBlock set a specific nextArweaveBlock in the store from its index
func (k Keeper) SetNextArweaveBlock(ctx sdk.Context, nextArweaveBlock types.NextArweaveBlock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextArweaveBlockKeyPrefix))
	b := k.cdc.MustMarshal(&nextArweaveBlock)
	store.Set(types.NextArweaveBlockKey(
		nextArweaveBlock.GetHeightString(),
	), b)
}

// GetNextArweaveBlock returns a nextArweaveBlock from its index
func (k Keeper) GetNextArweaveBlock(
	ctx sdk.Context,
	height string,

) (val types.NextArweaveBlock, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextArweaveBlockKeyPrefix))

	b := store.Get(types.NextArweaveBlockKey(
		height,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNextArweaveBlock removes a nextArweaveBlock from the store
func (k Keeper) RemoveNextArweaveBlock(
	ctx sdk.Context,
	height string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextArweaveBlockKeyPrefix))
	store.Delete(types.NextArweaveBlockKey(
		height,
	))
}

// GetAllNextArweaveBlock returns all nextArweaveBlock
func (k Keeper) GetAllNextArweaveBlock(ctx sdk.Context) (list []types.NextArweaveBlock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextArweaveBlockKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NextArweaveBlock
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
