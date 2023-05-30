package ante

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// TODO add tests for other scenarios when support for Arweave/EVM keys is added
func TestVerifySignaturesPublicKeyMismatch(t *testing.T) {
	_, pubKey, _ := testdata.KeyTestPubAddr()
	dataItem := exampleDataItem()

	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	creator, _ := sdk.AccAddressFromBech32(dataItem.Creator)
	acc := app.AccountKeeper.NewAccountWithAddress(ctx, creator)
	app.AccountKeeper.SetAccount(ctx, acc)

	sigV2 := signing.SignatureV2{
		PubKey: pubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: dataItem.DataItem.Signature,
		},
		Sequence: acc.GetSequence(),
	}

	txBuilder := newTxBuilder()
	txBuilder.SetMsgs(&dataItem)
	txBuilder.SetSignatures(sigV2)
	tx := txBuilder.GetTx()

	err := verifySignatures(ctx, app.AccountKeeper, tx, &dataItem)

	require.ErrorIs(t, err, types.ErrPublicKeyMismatch)
}
