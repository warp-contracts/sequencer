package ante

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestArweaveBlockTx(t *testing.T) {
	block := test.ArweaveBlock()
	tx := createTxWithMsgs(t, &block)

	err := verifyArweaveBlockTx(tx)

	require.NoError(t, err)
}

func TestArweaveBlockTxWithAnotherMessageAfter(t *testing.T) {
	block := test.ArweaveBlock()
	dataItem := test.ArweaveL2Interaction(t)
	tx := createTxWithMsgs(t, &block, &dataItem)

	err := verifyArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestArweaveBlockTxWithAnotherMessageBefore(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	block := test.ArweaveBlock()
	tx := createTxWithMsgs(t, &dataItem, &block)

	err := verifyArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestArweaveBlockTxWithoutBlock(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	tx := createTxWithMsgs(t, &dataItem)

	err := verifyArweaveBlockTx(tx)

	require.NoError(t, err)
}

func TestArweaveBlockTxWithoutMsgs(t *testing.T) {
	tx := createTxWithMsgs(t)

	err := verifyArweaveBlockTx(tx)

	require.NoError(t, err)
}
