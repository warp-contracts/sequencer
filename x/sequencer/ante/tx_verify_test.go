package ante

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestGetDataItemMsgOneDataItem(t *testing.T) {
	dataItem := exampleDataItem()
	txBuilder := newTxBuilder()
	txBuilder.SetMsgs(&dataItem)
	tx := txBuilder.GetTx()

	result, err := GetDataItemMsg(tx)

	require.NoError(t, err)
	require.Equal(t, &dataItem, result)
}

func TestGetDataItemMsgNoMsgs(t *testing.T) {
	txBuilder := newTxBuilder()
	tx := txBuilder.GetTx()

	result, err := GetDataItemMsg(tx)

	require.Nil(t, err)
	require.Nil(t, result)
}

func TestGetDataItemMsgTooManyDataItems(t *testing.T) {
	dataItem := exampleDataItem()
	txBuilder := newTxBuilder()
	txBuilder.SetMsgs(&dataItem, &dataItem)
	tx := txBuilder.GetTx()

	result, err := GetDataItemMsg(tx)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestGetDataItemMsgDataItemBeforeMsg(t *testing.T) {
	dataItem := exampleDataItem()
	msg := testdata.NewTestMsg(sdk.AccAddress(dataItem.Creator))
	txBuilder := newTxBuilder()
	txBuilder.SetMsgs(&dataItem, msg)
	tx := txBuilder.GetTx()

	result, err := GetDataItemMsg(tx)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestGetDataItemMsgDataItemAfterMsg(t *testing.T) {
	dataItem := exampleDataItem()
	msg := testdata.NewTestMsg(sdk.AccAddress(dataItem.Creator))
	txBuilder := newTxBuilder()
	txBuilder.SetMsgs(msg, &dataItem)
	tx := txBuilder.GetTx()

	result, err := GetDataItemMsg(tx)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}
