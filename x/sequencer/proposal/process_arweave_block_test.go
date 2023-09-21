package proposal

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func ctxHandlerAndLogger(t *testing.T, lastBlock *types.LastArweaveBlock, nextBlock *types.ArweaveBlockInfo) (sdk.Context, *processProposalHandler, *LoggerMock) {
	keeper, ctx := keepertest.SequencerKeeper(t)

	if lastBlock != nil {
		keeper.SetLastArweaveBlock(ctx, *lastBlock)
	}

	blockHeader := ctx.BlockHeader()
	blockHeader.Time = time.Unix(1692357017, 0)
	blockHeader.Height = 123

	controller := controller.MockArweaveBlocksController(nextBlock)
	logger := &LoggerMock{}

	handler := &processProposalHandler{keeper: keeper, controller: controller, logger: logger}
	return ctx.WithBlockHeader(blockHeader), handler, logger
}

func TestValidateIndex(t *testing.T) {
	_, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	result := handler.validateIndex(0)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestValidateIndexNotFirst(t *testing.T) {
	_, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	result := handler.validateIndex(1)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: Arweave block must be in the first transaction in the sequencer block")
}

func TestValidateArweaveBlockTx(t *testing.T) {
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block)
	_, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	result := handler.validateArweaveBlockTx(tx)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestValidateArweaveBlockTxNoMsgs(t *testing.T) {
	tx := test.CreateTxWithMsgs(t)
	_, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	result := handler.validateArweaveBlockTx(tx)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: transaction with Arweave block must have exactly one message")
}

func TestValidateArweaveBlockTxToManyMsgs(t *testing.T) {
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block, &block)
	_, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	result := handler.validateArweaveBlockTx(tx)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: transaction with Arweave block must have exactly one message")
}

func TestValidateArweaveBlockMsg(t *testing.T) {
	block := test.ArweaveBlock()
	ctx, handler, logger := ctxHandlerAndLogger(t, nil, block.BlockInfo)

	result := handler.validateArweaveBlockMsg(ctx, &block)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestValidateArweaveBlockMsgWithoutHoursDelay(t *testing.T) {
	ctx, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	block := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692357016,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}

	result := handler.validateArweaveBlockMsg(ctx, block)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: Arweave block should be one hour older than the sequencer block")
}

func TestValidateArweaveBlockMsgWithoutNextHeight(t *testing.T) {
	block := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692353410,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil)

	result := handler.validateArweaveBlockMsg(ctx, block)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: new height of the Arweave block is not the next value compared to the previous height")
}

func TestValidateArweaveBlockMsgWithoutLaterTimestamp(t *testing.T) {
	oldBlock := &types.ArweaveBlockInfo{
		Height:    1431216,
		Timestamp: 1692353410,
		Hash:      test.ExampleArweaveBlockHash,
	}
	newBlock := &types.ArweaveBlockInfo{
		Height:    1431217,
		Timestamp: 1692353410,
		Hash:      test.ExampleArweaveBlockHash,
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: oldBlock,
	}, nil)

	result := handler.validateArweaveBlockMsg(ctx, &types.MsgArweaveBlock{
		BlockInfo: newBlock,
	})

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: timestamp of the Arweave block is not later than the previous one")
}

func TestValidateArweaveBlockMsgWithoutNextArweaveBlock(t *testing.T) {
	ctx, handler, logger := ctxHandlerAndLogger(t, nil, nil)
	block := test.ArweaveBlock()

	result := handler.validateArweaveBlockMsg(ctx, &block)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: the Validator did not fetch the Arweave block with given height")
}

func TestValidateArweaveBlockMsgTimestampMismatchWithNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	nextBlockInfo := *block.BlockInfo
	nextBlockInfo.Timestamp += 1
	ctx, handler, logger := ctxHandlerAndLogger(t, nil, &nextBlockInfo)

	result := handler.validateArweaveBlockMsg(ctx, &block)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: timestamp of the Arweave block does not match the timestamp of the block downloaded by the Validator")
}

func TestValidateArweaveBlockMsgServerHashMismatchWithNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	nextBlockInfo := *block.BlockInfo
	nextBlockInfo.Hash = "abc"
	ctx, handler, logger := ctxHandlerAndLogger(t, nil, &nextBlockInfo)

	result := handler.validateArweaveBlockMsg(ctx, &block)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: hash of the Arweave block does not match the hash of the block downloaded by the Validator")
}

func TestArweaveBlockIsNotMissingGenesis(t *testing.T) {
	ctx, handler, logger := ctxHandlerAndLogger(t, nil, nil)
	ctx = ctx.WithBlockHeight(0)

	result := handler.checkArweaveBlockIsNotMissing(ctx, 0)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestArweaveBlockIsNotMissingNotFirst(t *testing.T) {
	ctx, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	result := handler.checkArweaveBlockIsNotMissing(ctx, 1)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestArweaveBlockIsNotMissingNotNext(t *testing.T) {
	block := test.ArweaveBlock()
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil)

	result := handler.checkArweaveBlockIsNotMissing(ctx, 0)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestArweaveBlockIsNotMissing(t *testing.T) {
	blockInfo := test.ArweaveBlock().BlockInfo
	blockInfo.Timestamp += 1692356017
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: blockInfo,
	}, blockInfo)

	result := handler.checkArweaveBlockIsNotMissing(ctx, 0)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestArweaveBlockIsMissing(t *testing.T) {
	blockInfo := test.ArweaveBlock().BlockInfo
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: blockInfo,
	}, blockInfo)

	result := handler.checkArweaveBlockIsNotMissing(ctx, 0)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: first transaction of the block should contain a transaction with the Arweave block")
}

func TestCheckTransactionsLengthMismatch(t *testing.T) {
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", SortKey: "123"}}
	block := &types.MsgArweaveBlock{
		BlockInfo: test.ArweaveBlock().BlockInfo,
	}
	_, handler, logger := ctxHandlerAndLogger(t, nil, nil)

	result := handler.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: incorrect number of transactions in the Arweave block")
}

func TestCheckTransactionsIdMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithLastSortKey{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "1234", SortKey: "1,2,3"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil)

	handler.initSortKeyForBlock(ctx)
	result := handler.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: transaction id is not as expected")
}

func TestCheckTransactionsContractMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithLastSortKey{{Transaction: &types.ArweaveTransaction{Contract: "abcd", Id: "123", SortKey: "1,2,3"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil)

	handler.initSortKeyForBlock(ctx)
	result := handler.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: the contract of the transaction does not match the expected one")
}

func TestCheckTransactionsSortKeyMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithLastSortKey{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3,4"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil)

	handler.initSortKeyForBlock(ctx)
	result := handler.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: transaction sort key is not as expected")
}

func TestCheckTransactionsLastSortKeyMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithLastSortKey{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil)
	handler.keeper.SetLastSortKey(ctx, types.LastSortKey{Contract: "abc", SortKey: "1,1,1"})

	handler.initSortKeyForBlock(ctx)
	result := handler.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid last sort key")
}

func TestCheckTransactions(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithLastSortKey{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3"}, LastSortKey: "1,1,1"}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil)
	handler.keeper.SetLastSortKey(ctx, types.LastSortKey{Contract: "abc", SortKey: "1,1,1"})

	handler.initSortKeyForBlock(ctx)
	result := handler.checkTransactions(block, expectedTxs)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}