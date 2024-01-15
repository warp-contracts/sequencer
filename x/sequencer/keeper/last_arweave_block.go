package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// SetLastArweaveBlock set lastArweaveBlock in the store
func (k Keeper) SetLastArweaveBlock(ctx context.Context, lastArweaveBlock types.LastArweaveBlock) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LastArweaveBlockKey))
	b := k.cdc.MustMarshal(&lastArweaveBlock)
	store.Set([]byte{0}, b)
}

// GetLastArweaveBlock returns lastArweaveBlock
func (k Keeper) GetLastArweaveBlock(ctx context.Context) (val types.LastArweaveBlock, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LastArweaveBlockKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) MustGetLastArweaveBlock(ctx context.Context) types.LastArweaveBlock {
	lastArweaveBlock, found := k.GetLastArweaveBlock(ctx)
	if !found {
		panic("LastArweaveBlock must be set")
	}
	return lastArweaveBlock
}

// RemoveLastArweaveBlock removes lastArweaveBlock from the store
func (k Keeper) RemoveLastArweaveBlock(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LastArweaveBlockKey))
	store.Delete([]byte{0})
}
