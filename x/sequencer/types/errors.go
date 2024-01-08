package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/sequencer module sentinel errors
var (
	ErrInvalidMessagesNumber       = errors.Register(ModuleName, 1100, "invalid number of messages")
	ErrNotEmptyMemo                = errors.Register(ModuleName, 1101, "not empty memo")
	ErrNonZeroTimeoutHeight        = errors.Register(ModuleName, 1102, "non-zero timeout height")
	ErrHasExtensionOptions         = errors.Register(ModuleName, 1103, "has extension options")
	ErrNotSingleSignature          = errors.Register(ModuleName, 1104, "not single signature")
	ErrInvalidSignMode             = errors.Register(ModuleName, 1105, "invalid sign mode")
	ErrNotEmptySignature           = errors.Register(ModuleName, 1106, "not empty signature")
	ErrTooManySigners              = errors.Register(ModuleName, 1107, "too many signers")
	ErrPublicKeyMismatch           = errors.Register(ModuleName, 1108, "public key mismatch")
	ErrNoSequencerNonceTag         = errors.Register(ModuleName, 1109, "no sequencer nonce tag")
	ErrNoContractTag               = errors.Register(ModuleName, 1110, "no contract tag")
	ErrSequencerNonceMismatch      = errors.Register(ModuleName, 1111, "sequencer nonce mismatch")
	ErrNonZeroGas                  = errors.Register(ModuleName, 1112, "non-zero gas")
	ErrNonZeroFee                  = errors.Register(ModuleName, 1113, "non-zero fee")
	ErrNotEmptyFeePayer            = errors.Register(ModuleName, 1114, "not empty fee payer")
	ErrNotEmptyFeeGranter          = errors.Register(ModuleName, 1115, "not empty fee granter")
	ErrInvalidSigner               = errors.Register(ModuleName, 1116, "expected gov account as only signer for proposal message")
	ErrBadArweaveHeight            = errors.Register(ModuleName, 1117, "invalid arweave block height")
	ErrBadArweaveTimestamp         = errors.Register(ModuleName, 1118, "invalid arweave block timestamp")
	ErrBadArweaveHash              = errors.Register(ModuleName, 1119, "invalid arweave block hash")
	ErrArweaveBlockNotFromProposer = errors.Register(ModuleName, 1120, "arweave block not from proposer")
	ErrArweaveBlockNotOldEnough    = errors.Register(ModuleName, 1121, "arweave block not old enough")
	ErrUnknownArweaveBlock         = errors.Register(ModuleName, 1122, "unknown arweave block")
	ErrArweaveBlockMissing         = errors.Register(ModuleName, 1123, "arweave block is missing")
	ErrDataItemVerification        = errors.Register(ModuleName, 1124, "data item verification error")
	ErrDataItemAlreadyInBlock      = errors.Register(ModuleName, 1125, "data item already in block")
	ErrInvalidSortKey              = errors.Register(ModuleName, 1126, "invalid sort key")
	ErrInvalidPrevSortKey          = errors.Register(ModuleName, 1127, "invalid prev sort key")
	ErrInvalidRandomValue          = errors.Register(ModuleName, 1128, "invalid random value")
	ErrInvalidTxIndex              = errors.Register(ModuleName, 1129, "invalid transaction index")
	ErrInvalidTxNumber             = errors.Register(ModuleName, 1130, "invalid number of transactions")
	ErrTxIdMismatch                = errors.Register(ModuleName, 1131, "transaction id mismatch")
	ErrTxContractMismatch          = errors.Register(ModuleName, 1132, "transaction contract mismatch")
)
