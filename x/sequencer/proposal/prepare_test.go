package proposal

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/types/module/testutil"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
)

// Checks if the method for calculating the transaction size is consistent with the Protobuf size
func TestProtoTxSize(t *testing.T) {
	txEncoder := testutil.MakeTestEncodingConfig().TxConfig.TxEncoder()

	block := test.ArweaveBlock()
	arL2 := test.ArweaveL2Interaction(t)
	ethL2 := test.EthereumL2Interaction(t)

	tx1, err1 := txEncoder(test.CreateTxWithMsgs(t, &block))
	tx2, err2 := txEncoder(test.CreateTxWithMsgs(t, &arL2))
	tx3, err3 := txEncoder(test.CreateTxWithMsgs(t, &ethL2))

	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NoError(t, err3)

	data := types.Data{
		Txs: [][]byte{tx1, tx2, tx3},
	}
	protoSize := int64(data.Size())
	estimatedSize := protoTxSize(tx1) + protoTxSize(tx2) + protoTxSize(tx3)

	require.Equal(t, protoSize, estimatedSize)
}
