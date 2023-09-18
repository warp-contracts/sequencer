package ante

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestGetDataItemMsgOneDataItem(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	tx := test.CreateTxWithMsgs(t, &dataItem)

	result, err := GetL2Interaction(tx)

	require.NoError(t, err)
	require.Equal(t, &dataItem, result)
}

func TestGetDataItemMsgNoMsgs(t *testing.T) {
	tx := test.CreateTxWithMsgs(t)

	result, err := GetL2Interaction(tx)

	require.Nil(t, err)
	require.Nil(t, result)
}

func TestGetDataItemMsgTooManyDataItems(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	tx := test.CreateTxWithMsgs(t, &dataItem, &dataItem)

	result, err := GetL2Interaction(tx)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestGetDataItemMsgDataItemBeforeMsg(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	msg := testdata.NewTestMsg(dataItem.GetCreator())
	tx := test.CreateTxWithMsgs(t, &dataItem, msg)

	result, err := GetL2Interaction(tx)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestGetDataItemMsgDataItemAfterMsg(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	msg := testdata.NewTestMsg(dataItem.GetCreator())
	tx := test.CreateTxWithMsgs(t, msg, &dataItem)

	result, err := GetL2Interaction(tx)

	require.Nil(t, result)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}
