package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// SetLastArweaveBlock set lastArweaveBlock in the store
func (k Keeper) SetLastArweaveBlock(ctx sdk.Context, lastArweaveBlock types.ArweaveBlockInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastArweaveBlockKey))
	b := k.cdc.MustMarshal(&lastArweaveBlock)
	store.Set([]byte{0}, b)
}

// GetLastArweaveBlock returns lastArweaveBlock
func (k Keeper) GetLastArweaveBlock(ctx sdk.Context) (val types.ArweaveBlockInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastArweaveBlockKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) MustGetLastArweaveBlock(ctx sdk.Context) types.ArweaveBlockInfo {
	blockInfo, found := k.GetLastArweaveBlock(ctx)
	if !found {
		panic("LastArweaveBlock must be set")
	}
	return blockInfo
}

// RemoveLastArweaveBlock removes lastArweaveBlock from the store
func (k Keeper) RemoveLastArweaveBlock(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastArweaveBlockKey))
	store.Delete([]byte{0})
}
