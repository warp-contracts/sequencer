package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sequencer module sentinel errors
var (
	ErrToManyDataItems = sdkerrors.Register(ModuleName, 1100, "to many data items")
	ErrNotEmptyMemo = sdkerrors.Register(ModuleName, 1101, "not empty memo")
	ErrNonZeroTimeoutHeight = sdkerrors.Register(ModuleName, 1102, "non-zero timeout height")
	ErrHasExtensionOptions = sdkerrors.Register(ModuleName, 1103, "has extension options")
	ErrToManySignatures = sdkerrors.Register(ModuleName, 1104, "to many signatures")
	ErrSignatureMismatch = sdkerrors.Register(ModuleName, 1105, "signature mismatch")
	ErrToManySigners = sdkerrors.Register(ModuleName, 1106, "to many signers")
	ErrPublicKeyMismatch = sdkerrors.Register(ModuleName, 1107, "public key mismatch")
	ErrNoSequencerNonceTag = sdkerrors.Register(ModuleName, 1108, "no sequencer nonce tag")
	ErrSequencerNonceMismatch = sdkerrors.Register(ModuleName, 1109, "sequencer nonce mismatch")
	ErrNonZeroGas = sdkerrors.Register(ModuleName, 1110, "non-zero gas")
	ErrNonZeroFee = sdkerrors.Register(ModuleName, 1111, "non-zero fee")
	ErrNotEmptyFeePayer = sdkerrors.Register(ModuleName, 1112, "not empty fee payer")
	ErrNotEmptyFeeGranter = sdkerrors.Register(ModuleName, 1113, "not empty fee granter")
	ErrNotEmptyTip = sdkerrors.Register(ModuleName, 1114, "not empty tip")
)
