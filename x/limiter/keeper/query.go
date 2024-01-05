package keeper

import (
	"github.com/warp-contracts/sequencer/x/limiter/types"
)

var _ types.QueryServer = Keeper{}
