package ante

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestArweaveBlockTx(t *testing.T) {
	blockInfo := test.ArweaveBlockInfo()
	interaction := test.ArweaveL1Interaction(t)
	tx := createTxWithMsgs(t, &blockInfo, &interaction)

	err := verifyArweaveBlockTx(tx)

	require.NoError(t, err)
}

func TestArweaveBlockTxNoBlockInfo(t *testing.T) {
	interaction := test.ArweaveL1Interaction(t)
	tx := createTxWithMsgs(t, &interaction)

	err := verifyArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrInvalidArweaveBlockTx)
}

func TestArweaveBlockTxOnlyBlockInfo(t *testing.T) {
	blockInfo := test.ArweaveBlockInfo()
	tx := createTxWithMsgs(t, &blockInfo)

	err := verifyArweaveBlockTx(tx)

	require.NoError(t, err)
}

func TestArweaveBlockTxBlockInfoAfterL1Interaction(t *testing.T) {
	interaction := test.ArweaveL1Interaction(t)
	blockInfo := test.ArweaveBlockInfo()
	tx := createTxWithMsgs(t, &interaction, &blockInfo)

	err := verifyArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrInvalidArweaveBlockTx)
}

func TestArweaveBlockTxBlockInfoAfterL2Interaction(t *testing.T) {
	interaction := test.ArweaveL2Interaction(t)
	blockInfo := test.ArweaveBlockInfo()
	tx := createTxWithMsgs(t, &interaction, &blockInfo)

	err := verifyArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrInvalidArweaveBlockTx)
}

func TestArweaveBlockTxWithL2Interaction(t *testing.T) {
	blockInfo := test.ArweaveBlockInfo()
	l1Interaction := test.ArweaveL1Interaction(t)
	l2Interaction := test.ArweaveL2Interaction(t)
	tx := createTxWithMsgs(t, &blockInfo, &l1Interaction, &l2Interaction)

	err := verifyArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrInvalidArweaveBlockTx)
}
