package ante

import (
	"time"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestArweaveBlockTxVerify(t *testing.T) {
	block := test.ArweaveBlock()
	tx := createTxWithMsgs(t, &block)

	foundTx, err := verifyArweaveBlockTx(tx)

	require.Equal(t, true, foundTx)
	require.NoError(t, err)
}

func TestArweaveBlockTxVerifyWithAnotherMessageAfter(t *testing.T) {
	block := test.ArweaveBlock()
	dataItem := test.ArweaveL2Interaction(t)
	tx := createTxWithMsgs(t, &block, &dataItem)

	foundTx, err := verifyArweaveBlockTx(tx)

	require.Equal(t, true, foundTx)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestArweaveBlockTxVerifyWithAnotherMessageBefore(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	block := test.ArweaveBlock()
	tx := createTxWithMsgs(t, &dataItem, &block)

	foundTx, err := verifyArweaveBlockTx(tx)

	require.Equal(t, true, foundTx)
	require.ErrorIs(t, err, types.ErrTooManyMessages)
}

func TestArweaveBlockTxVerifyWithoutBlock(t *testing.T) {
	dataItem := test.ArweaveL2Interaction(t)
	tx := createTxWithMsgs(t, &dataItem)

	foundTx, err := verifyArweaveBlockTx(tx)

	require.Equal(t, false, foundTx)
	require.NoError(t, err)
}

func TestArweaveBlockTxVerifyWithoutMsgs(t *testing.T) {
	tx := createTxWithMsgs(t)

	foundTx, err := verifyArweaveBlockTx(tx)

	require.Equal(t, false, foundTx)
	require.NoError(t, err)
}

func keeperAndContext(t *testing.T, blockTimestamp int64, lastTimestamp uint64, nextTimestamp uint64) (*keeper.Keeper, sdk.Context) {
	k, ctx := keepertest.SequencerKeeper(t)
	if lastTimestamp > 0 {
		k.SetLastArweaveBlock(ctx, types.ArweaveBlockInfo{
			Height: 1,
			Timestamp: lastTimestamp,
		})	
	}
	if nextTimestamp > 0 {
		k.SetNextArweaveBlock(ctx, types.NextArweaveBlock{
			BlockInfo: &types.ArweaveBlockInfo{
				Height: 2,
				Timestamp: nextTimestamp,
			},
		})
	}
	blockHeader := ctx.BlockHeader()
	blockHeader.Time = time.Unix(blockTimestamp, 0)
	return k, ctx.WithBlockHeader(blockHeader)
}

func TestArweaveBlockTxNoNeedArweaveTx(t *testing.T) {
	k, ctx := keeperAndContext(t, 200, 100, 300)

	err := shouldBlockContainArweaveTx(ctx, k)

	require.NoError(t, err)
}

func TestArweaveBlockTxWithoutNextArweaveBlock(t *testing.T) {
	k, ctx := keeperAndContext(t, 200, 100, 0)

	err := shouldBlockContainArweaveTx(ctx, k)

	require.NoError(t, err)
}

func TestArweaveBlockTxShouldContainArweaveBlock(t *testing.T) {
	k, ctx := keeperAndContext(t, 10000, 100, 300)

	err := shouldBlockContainArweaveTx(ctx, k)

	require.ErrorIs(t, err, types.ErrNoArweaveBlockTx)
}
