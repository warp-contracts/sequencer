package ante

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func appAndCtx(t *testing.T) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	return app, ctx
}

func addCreatorAccount(app *simapp.SimApp, ctx sdk.Context, dataItem types.MsgDataItem) authtypes.AccountI {
	creator, _ := sdk.AccAddressFromBech32(dataItem.Creator)
	acc := app.AccountKeeper.NewAccountWithAddress(ctx, creator)
	app.AccountKeeper.SetAccount(ctx, acc)
	return acc
}

func createSignature(sequence uint64, data signing.SignatureData) (signing.SignatureV2) {
	_, pubKey, _ := testdata.KeyTestPubAddr()
	return signing.SignatureV2{
		PubKey: pubKey,
		Data: data,
		Sequence: sequence,
	}
}

func createEmptySignature(sequence uint64) (signing.SignatureV2) {
	data := &signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	return createSignature(sequence, data)
}

func createTxWithSignatures(t *testing.T, dataItem types.MsgDataItem, signatures ...signing.SignatureV2) (authsigning.Tx) {
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
	sig := createEmptySignature(acc.GetSequence())
	tx := createTxWithSignatures(t, dataItem, sig, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrNotSingleSignature)
}

func TestVerifySignaturesNoSignerAccount(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	sig := createEmptySignature(0)
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
	sig := createSignature(acc.GetSequence(), sigData)
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrNotEmptySignature)
}

func TestVerifySignaturesMultiSignature(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	sigData := &signing.MultiSignatureData{}
	sig := createSignature(acc.GetSequence(), sigData)
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrTooManySigners)
}

func TestVerifySignaturesPublicKeyMismatch(t *testing.T) {
	app, ctx := appAndCtx(t)
	dataItem := exampleDataItem(t)
	acc := addCreatorAccount(app, ctx, dataItem)
	sig := createEmptySignature(acc.GetSequence())
	tx := createTxWithSignatures(t, dataItem, sig)

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrPublicKeyMismatch)
}

// TODO tests for nonce when support for Arweave/EVM keys is added