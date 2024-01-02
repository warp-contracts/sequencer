package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Each block height has a separate prefix
// Each limiter kind has a separate prefix
func (k *Keeper) getStore(ctx sdk.Context, blockHeight int64, limiterIndex int) prefix.Store {
	return prefix.NewStore(prefix.NewStore(ctx.KVStore(k.storeKey), []byte(strconv.FormatInt(blockHeight, 10))), []byte(strconv.FormatInt(int64(limiterIndex), 10)))
}

func (k *Keeper) GetCount(ctx sdk.Context, limiterIndex int, key []byte) int64 {
	return k.controller.Cache.GetCount(limiterIndex, string(key))
}

func (k *Keeper) SetCurrentBlockHeight(ctx sdk.Context, blockHeight int64) {
	// Setup filling in the cache
	if k.lastInitHeight <= 0 {
		k.lastInitHeight = blockHeight + k.numCachedBlocks
		return
	}

	if blockHeight == k.currentBlockHeight+1 {
		// This handles most of the cases
		k.currentBlockHeight = blockHeight
		return
	}

	k.Logger(ctx).Info("Cleaning up limiter cache", "currentBlockHeight", k.currentBlockHeight, "blockHeight", blockHeight)
	k.currentBlockHeight = blockHeight
	k.Clean(ctx, blockHeight)
}

func (k *Keeper) Inc(ctx sdk.Context, limiterIndex int, key []byte) {
	store := k.getStore(ctx, k.currentBlockHeight, limiterIndex)
	value := store.Get(key)
	if value == nil {
		// Initial value
		store.Set(key, []byte("1"))
		k.controller.Cache.Set(limiterIndex, string(key), 1)
		return
	}

	// Parse existing value
	i, err := strconv.ParseInt(string(value), 10, 64)
	if err != nil {
		panic(err)
	}
	i += 1

	// Set value
	store.Set(key, []byte(strconv.FormatInt(i, 10)))

	// Update cached counters
	k.controller.Cache.Set(limiterIndex, string(key), i)
}

/*
                     ┌─────────────────────────────┐
                     │                             │
                     │           CACHE             │
                     │                             │                                        h
───┬─────┬─────┬─────┼─────────────────────────────┼─────┬─────┬─────┬─────┬─────┬────────────►
   │     │     │     │                             │     │     │     │     │     │
                   start                         finish  │
                                                         │

                                                 currentBlockHeight
*/
// Goes through each block height and deletes all the keys
// Subtracts the value from the cached counters to keep them up to date
func (k *Keeper) Clean(ctx sdk.Context, newFinish int64 /* new finish block height */) {
	if newFinish <= k.lastInitHeight {
		// It's still in the initial filling in phase
		return
	}

	// The last block height to keep in the cache
	newStart := newFinish - k.numCachedBlocks

	// Remove keys between old and new start heights
	// Most probably there will be only one iteration
	for h := k.start; h < newStart; h++ {

		// Iterate over all limiter kinds
		for limiterIdx := 0; limiterIdx < k.numLimiters; limiterIdx++ {
			store := k.getStore(ctx, h, limiterIdx)
			iter := store.Iterator(nil, nil)
			defer iter.Close()
			for ; iter.Valid(); iter.Next() {
				// Parse value
				i, err := strconv.ParseInt(string(iter.Value()), 10, 64)
				if err != nil {
					panic(err)
				}

				// Subtract value from the cached counter
				k.controller.Cache.Subtract(limiterIdx, string(iter.Key()), i)

				// Cleanup the kvstore
				store.Delete(iter.Key())
			}
		}
	}

	// Update cache range indices
	k.start = newStart
	k.finish = newFinish
}
