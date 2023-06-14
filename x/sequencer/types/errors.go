package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sequencer module sentinel errors
var (
	ErrTooManyMessages        = sdkerrors.Register(ModuleName, 1100, "too many messages")
	ErrNotEmptyMemo           = sdkerrors.Register(ModuleName, 1101, "not empty memo")
	ErrNonZeroTimeoutHeight   = sdkerrors.Register(ModuleName, 1102, "non-zero timeout height")
	ErrHasExtensionOptions    = sdkerrors.Register(ModuleName, 1103, "has extension options")
	ErrNotSingleSignature     = sdkerrors.Register(ModuleName, 1104, "not single signature")
	ErrInvalidSignMode        = sdkerrors.Register(ModuleName, 1105, "invalid sign mode")
	ErrNotEmptySignature      = sdkerrors.Register(ModuleName, 1106, "not empty signature")
	ErrTooManySigners         = sdkerrors.Register(ModuleName, 1107, "too many signers")
	ErrPublicKeyMismatch      = sdkerrors.Register(ModuleName, 1108, "public key mismatch")
	ErrNoSequencerNonceTag    = sdkerrors.Register(ModuleName, 1109, "no sequencer nonce tag")
	ErrSequencerNonceMismatch = sdkerrors.Register(ModuleName, 1110, "sequencer nonce mismatch")
	ErrNonZeroGas             = sdkerrors.Register(ModuleName, 1111, "non-zero gas")
	ErrNonZeroFee             = sdkerrors.Register(ModuleName, 1112, "non-zero fee")
	ErrNotEmptyFeePayer       = sdkerrors.Register(ModuleName, 1113, "not empty fee payer")
	ErrNotEmptyFeeGranter     = sdkerrors.Register(ModuleName, 1114, "not empty fee granter")
	ErrNotEmptyTip            = sdkerrors.Register(ModuleName, 1115, "not empty tip")
)
