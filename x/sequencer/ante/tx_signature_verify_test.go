package ante

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	keys "github.com/warp-contracts/sequencer/crypto/keys/arweave"
	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/utils/bundlr"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func appAndCtx(t *testing.T) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	return app, ctx
}

func addCreatorAccount(app *simapp.SimApp, ctx sdk.Context, dataItem types.MsgDataItem) authtypes.AccountI {
	acc := app.AccountKeeper.NewAccountWithAddress(ctx, dataItem.GetCreator())
	app.AccountKeeper.SetAccount(ctx, acc)
	return acc
}

func createSignature(dataItem types.MsgDataItem, sequence uint64, data signing.SignatureData) signing.SignatureV2 {
	pubKey := &keys.PubKey{Key: dataItem.DataItem.Owner}
	return signing.SignatureV2{
		PubKey:   pubKey,
		Sequence: sequence,
		Data:     data,
	}
}

var singleSignatureData = &signing.SingleSignatureData{
	SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
	Signature: nil,
}

func createEmptySignature(dataItem types.MsgDataItem, sequence uint64) signing.SignatureV2 {
	return createSignature(dataItem, sequence, singleSignatureData)
}

func createTxWithSignatures(t *testing.T, dataItem types.MsgDataItem, signatures ...signing.SignatureV2) authsigning.Tx {
	txBuilder := newTxBuilder()

	err := txBuilder.SetMsgs(&dataItem)
	require.NoError(t, err)

	err = txBuilder.SetSignatures(signatures...)
	require.NoError(t, err)

	return txBuilder.GetTx()
}

func TestVerifySignaturesNoSignatures(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	addCreatorAccount(app, ctx, dataItem)
	tx := createTxWithSignatures(t, dataItem)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrNotSingleSignature)
}

func TestVerifySignaturesTooManySignatures(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	sig := createEmptySignature(dataItem, acc.GetSequence())
	tx := createTxWithSignatures(t, dataItem, sig, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrNotSingleSignature)
}

func TestVerifySignaturesNoSignerAccount(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	sig := createEmptySignature(dataItem, 0)
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, sdkerrors.ErrUnknownAddress)
}

func TestVerifySignaturesNotEmptySignature(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	sigData := &signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: []byte("signature"),
	}
	sig := createSignature(dataItem, acc.GetSequence(), sigData)
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrNotEmptySignature)
}

func TestVerifySignaturesMultiSignature(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	sigData := &signing.MultiSignatureData{}
	sig := createSignature(dataItem, acc.GetSequence(), sigData)
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrTooManySigners)
}

func TestVerifySignaturesPublicKeyMismatch(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	_, pubKey, _ := testdata.KeyTestPubAddr()
	sig := 	signing.SignatureV2{
		PubKey:   pubKey,
		Sequence: acc.GetSequence(),
		Data:     singleSignatureData,
	}
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrPublicKeyMismatch)
}

func TestVerifySignaturesWrongSequence(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	sig := createEmptySignature(dataItem, acc.GetSequence() + 1)
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, sdkerrors.ErrWrongSequence)
}

func TestVerifySignaturesNoSequencerNonceTag(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	sig := createEmptySignature(dataItem, acc.GetSequence())
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrNoSequencerNonceTag)
}

func TestVerifySignaturesSequencerNonceMismatch(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t, bundlr.Tag{Name: "Sequencer-Nonce", Value: "1"})
	acc := addCreatorAccount(app, ctx, dataItem)
	sig := createEmptySignature(dataItem, acc.GetSequence())
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrSequencerNonceMismatch)
}

func TestVerifySignatures(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t, bundlr.Tag{Name: "Sequencer-Nonce", Value: "0"})
	acc := addCreatorAccount(app, ctx, dataItem)
	sig := createEmptySignature(dataItem, acc.GetSequence())
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.NoError(t, err)
}
