package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/warp-contracts/sequencer/x/limiter/controller"
	"github.com/warp-contracts/sequencer/x/limiter/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		// Current cache  heights
		start  int64
		finish int64

		// Last block height of the first block in the cache, used for the fill-in pahse
		lastInitHeight int64

		// Block height of the last block in the cache, used for getting the right kvstore
		currentBlockHeight int64

		// Number of blocks to keep in the cache
		numCachedBlocks int64

		// Cache of the limits
		cache []map[string]int64

		// Controller
		controller *controller.Controller
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	//  Number of limiters, indexed from 0
	numLimiters int,
	// Number of blocks to keep in the cache
	numCachedBlocks int64,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	// Setup async tasks
	ctrl, err := controller.NewController(log.NewNopLogger(), numLimiters)
	if err != nil {
		panic(err)
	}

	return &Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		memKey:          memKey,
		paramstore:      ps,
		cache:           make([]map[string]int64, numLimiters),
		lastInitHeight:  -1,
		numCachedBlocks: numCachedBlocks,
		controller:      ctrl,
	}
}

func (k Keeper) StopWait() {
	k.controller.StopWait()
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
