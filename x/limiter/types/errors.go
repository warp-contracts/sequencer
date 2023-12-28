package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/limiter module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
