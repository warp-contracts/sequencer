package ante

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
	"strconv"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func verifySignatures(ctx sdk.Context, ak sdkante.AccountKeeper, tx sdk.Tx, dataItem *types.MsgDataItem) error {
	sigTx, ok := tx.(signing.SigVerifiableTx)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "transaction is not of type SigVerifiableTx")
	}

	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return err
	}

	if err := verifySingleSignature(ctx, ak, sigs, dataItem); err != nil {
		return err
	}

	return nil
}

func verifySingleSignature(ctx sdk.Context, ak sdkante.AccountKeeper, sigs []txsigning.SignatureV2, dataItem *types.MsgDataItem) error {
	if len(sigs) != 1 {
		return sdkerrors.Wrap(types.ErrToManySignatures, "transaction with data item must contain exactly one signature")
	}

	sig := sigs[0]
	signer := dataItem.GetSigners()[0]
	acc, err := sdkante.GetSignerAcc(ctx, ak, signer)
	if err != nil {
		return err
	}

	switch sigData := sig.Data.(type) {
	case *txsigning.SingleSignatureData:
		if !bytes.Equal(sigData.Signature, dataItem.DataItem.Signature) {
			return sdkerrors.Wrap(types.ErrSignatureMismatch,
				"transaction with data item signature is different from data item signature")
		}
	case *txsigning.MultiSignatureData:
		return sdkerrors.Wrap(types.ErrToManySigners, "transaction with data item can only have one signer")
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

	tagSequence, err := getSequenceFromTags(dataItem)
	if err != nil {
		return err
	}

	if sig.Sequence != tagSequence {
		sdkerrors.Wrap(types.ErrSequencerNonceMismatch, "transaction sequence does not match nonce from data item tag")
	}

	return nil
}

func getSequenceFromTags(dataItem *types.MsgDataItem) (uint64, error) {
	const sequencerNonceTag = "Sequencer-Nonce"
	for _, tag := range dataItem.DataItem.Tags {
		if tag.Name == sequencerNonceTag {
			return strconv.ParseUint(tag.Value, 10, 64)
		}
	}
	return 0, sdkerrors.Wrapf(types.ErrNoSequencerNonceTag, "data item does not have \"%s\" tag", sequencerNonceTag)
}
