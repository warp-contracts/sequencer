package ante

import (
	"bytes"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Validation of the signature for a transaction with a DataItem.
// The transaction's signature must match the signature of the DataItem.
// Additionally, the nonce for the given sender is validated.
func verifySignaturesAndNonce(ctx sdk.Context, ak authkeeper.AccountKeeper, tx sdk.Tx, dataItem *types.MsgDataItem) error {
	sigTx, ok := tx.(signing.SigVerifiableTx)
	if !ok {
		return errors.Wrap(sdkerrors.ErrTxDecode, "transaction is not of type SigVerifiableTx")
	}

	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return err
	}

	if len(sigs) != 1 {
		return errors.Wrapf(types.ErrNotSingleSignature, "transaction with data item must contain exactly one signature, it has: %d", len(sigs))
	}

	sig := sigs[0]
	signer := dataItem.GetSigners()[0]

	if !ctx.IsReCheckTx() { // the signature does not need to be rechecked
		if err := verifySingleSignature(sig, signer, dataItem); err != nil {
			return err
		}
	}

	if err = verifyNonceAndIncreaseSequence(ctx, ak, sig, signer, dataItem); err != nil {
		return err
	}

	return nil
}


func verifySingleSignature(sig txsigning.SignatureV2, signer sdk.AccAddress, dataItem *types.MsgDataItem) error {
	switch sigData := sig.Data.(type) {
	case *txsigning.SingleSignatureData:
		if sigData.SignMode != txsigning.SignMode_SIGN_MODE_DIRECT {
			return errors.Wrap(types.ErrInvalidSignMode, "transaction with data item should have direct sign mode")
		}
		if len(sigData.Signature) > 0 {
			return errors.Wrap(types.ErrNotEmptySignature, "transaction with data item should have empty signature")
		}
	case *txsigning.MultiSignatureData:
		return errors.Wrap(types.ErrTooManySigners, "transaction with data item can only have one signer")
	}

	if !bytes.Equal(sig.PubKey.Address(), signer) {
		return errors.Wrap(types.ErrPublicKeyMismatch,
			"transaction public key address does not match message creator address")
	}

	if !bytes.Equal(sig.PubKey.Bytes(), dataItem.DataItem.Owner) {
		return errors.Wrap(types.ErrPublicKeyMismatch,
			"transaction public key does not match message public key")
	}

	return nil
}

func verifyNonceAndIncreaseSequence(ctx sdk.Context, ak authkeeper.AccountKeeper, sig txsigning.SignatureV2, signer sdk.AccAddress, dataItem *types.MsgDataItem) error {
	acc, err := getOrCreateAccount(ctx, ak, signer, dataItem)
	if err != nil {
		return err
	}

	if sig.Sequence != acc.GetSequence() {
		return errors.Wrapf(sdkerrors.ErrWrongSequence,
			"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
		)
	}

	tagNonce, err := dataItem.GetNonceFromTags()
	if err != nil {
		return err
	}

	if sig.Sequence != tagNonce {
		return errors.Wrap(types.ErrSequencerNonceMismatch, "transaction sequence does not match nonce from data item tag")
	}

	// increasing the account sequence
	if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
		return err
	}
	ak.SetAccount(ctx, acc)

	return nil
}

func getOrCreateAccount(ctx sdk.Context, ak authkeeper.AccountKeeper, addr sdk.AccAddress, dataItem *types.MsgDataItem) (authtypes.AccountI, error) {
	acc := ak.GetAccount(ctx, addr)

	if acc != nil {
		return acc, nil
	}

	pubKey, err := dataItem.GetPublicKey()
	if err != nil {
		return nil, err
	}

	acc = ak.NewAccountWithAddress(ctx, addr)

	err = acc.SetPubKey(pubKey)
	if err != nil {
		return nil, err
	}

	ak.SetAccount(ctx, acc)
	return acc, nil
}
