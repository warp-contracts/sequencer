package ante

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Validation of the signature for a transaction with a DataItem.
// The transaction's signature must match the signature of the DataItem.
// Additionally, the nonce for the given sender is validated.
func verifySignatures(ctx sdk.Context, ak sdkante.AccountKeeper, tx sdk.Tx, dataItem *types.MsgDataItem) error {
	sigTx, ok := tx.(signing.SigVerifiableTx)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "transaction is not of type SigVerifiableTx")
	}

	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return err
	}

	if len(sigs) != 1 {
		return sdkerrors.Wrapf(types.ErrNotSingleSignature, "transaction with data item must contain exactly one signature, it has: %d", len(sigs))
	}

	sig := sigs[0]
	signer := dataItem.GetSigners()[0]
	acc, err := sdkante.GetSignerAcc(ctx, ak, signer)
	if err != nil {
		return err
	}

	if err := verifySingleSignature(sig, signer, acc, dataItem); err != nil {
		return err
	}

	return nil
}

func verifySingleSignature(sig txsigning.SignatureV2, signer sdk.AccAddress, acc authtypes.AccountI, dataItem *types.MsgDataItem) error {
	switch sigData := sig.Data.(type) {
	case *txsigning.SingleSignatureData:
		if sigData.Signature != nil {
			return sdkerrors.Wrap(types.ErrNotEmptySignature, "transaction with data item should have empty signature")
		}
	case *txsigning.MultiSignatureData:
		return sdkerrors.Wrap(types.ErrTooManySigners, "transaction with data item can only have one signer")
	}

	if !bytes.Equal(sig.PubKey.Address(), signer) {
		return sdkerrors.Wrap(types.ErrPublicKeyMismatch,
			"transaction public key address does not match message creator address")
	}

	if !bytes.Equal(sig.PubKey.Bytes(), dataItem.DataItem.Owner) {
		return sdkerrors.Wrap(types.ErrPublicKeyMismatch,
			"transaction public key does not match message public key")
	}

	if err := verifyNonce(acc, sig, signer, dataItem); err != nil {
		return err
	}

	return nil
}

func verifyNonce(acc authtypes.AccountI, sig txsigning.SignatureV2, signer sdk.AccAddress, dataItem *types.MsgDataItem) error {
	if sig.Sequence != acc.GetSequence() {
		return sdkerrors.Wrapf(sdkerrors.ErrWrongSequence,
			"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
		)
	}

	tagSequence, err := dataItem.GetSequenceFromTags()
	if err != nil {
		return err
	}

	if sig.Sequence != tagSequence {
		return sdkerrors.Wrap(types.ErrSequencerNonceMismatch, "transaction sequence does not match nonce from data item tag")
	}

	return nil
}
