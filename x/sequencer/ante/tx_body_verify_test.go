package ante

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/stretchr/testify/require"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
	"testing"
)

func newTxBuilder() authtx.ExtensionOptionsTxBuilder {
	return simapp.MakeTestEncodingConfig().TxConfig.NewTxBuilder().(authtx.ExtensionOptionsTxBuilder)
}

func newAnyValue() *codectypes.Any {
	any, err := codectypes.NewAnyWithValue(testdata.NewTestMsg())
	if err != nil {
		panic(err)
	}
	return any
}

func TestVerifyTxBody(t *testing.T) {
	tx := newTxBuilder().GetTx()

	err := verifyTxBody(tx)

	require.NoError(t, err)
}

func TestVerifyTxBodyWithMemo(t *testing.T) {
	txBuilder := newTxBuilder()
	txBuilder.SetMemo("not empty memo")
	tx := txBuilder.GetTx()

	err := verifyTxBody(tx)

	require.ErrorIs(t, err, types.ErrNotEmptyMemo)
}

func TestVerifyTxBodyWithTimeoutHeight(t *testing.T) {
	txBuilder := newTxBuilder()
	txBuilder.SetTimeoutHeight(123)
	tx := txBuilder.GetTx()

	err := verifyTxBody(tx)

	require.ErrorIs(t, err, types.ErrNonZeroTimeoutHeight)
}

func TestVerifyTxBodyWithExtensionOptions(t *testing.T) {
	txBuilder := newTxBuilder()
	txBuilder.SetExtensionOptions(newAnyValue())
	tx := txBuilder.GetTx()

	err := verifyTxBody(tx)

	require.ErrorIs(t, err, types.ErrHasExtensionOptions)
}

func TestVerifyTxBodyWithNonCriticalExtensionOptions(t *testing.T) {
	txBuilder := newTxBuilder()
	txBuilder.SetNonCriticalExtensionOptions(newAnyValue())
	tx := txBuilder.GetTx()

	err := verifyTxBody(tx)

	require.ErrorIs(t, err, types.ErrHasExtensionOptions)
}
