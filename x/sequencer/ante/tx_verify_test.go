package ante

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
	"github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/bundlr"
)

func newTxBuilder() client.TxBuilder {
	return simapp.MakeTestEncodingConfig().TxConfig.NewTxBuilder()
}

func exampleDataItem() types.MsgDataItem {
	dataItem := bundlr.BundleItem{
		SignatureType: 1,
		Signature:     arweave.Base64String("signature"),
	}

	return types.MsgDataItem{
		Creator:  "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
		DataItem: dataItem,
	}
}

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
