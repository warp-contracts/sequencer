package proposal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func validatorAndLogger(t *testing.T, lastBlock *types.LastArweaveBlock, nextBlock *types.ArweaveBlockInfo, lastSortKey *types.LastSortKey) (*TxValidator, *LoggerMock) {
	keeper, ctx := keepertest.SequencerKeeper(t)

	if lastBlock == nil {
		keeper.SetLastArweaveBlock(ctx, types.LastArweaveBlock{ArweaveBlock: &types.ArweaveBlockInfo{Height: 1431215}})
	} else {
		keeper.SetLastArweaveBlock(ctx, *lastBlock)
	}

	if lastSortKey != nil {
		keeper.SetLastSortKey(ctx, *lastSortKey)
	}

	blockHeader := ctx.BlockHeader()
	blockHeader.Time = time.Unix(1692357017, 0)
	blockHeader.Height = 123

	controller := controller.MockArweaveBlocksController(nextBlock)
	logger := &LoggerMock{}

	validator := newTxValidator(ctx.WithBlockHeader(blockHeader), keeper, controller, logger)
	return validator, logger
}

func TestValidateLastSortKeysMismatch(t *testing.T) {
	txs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3"}, LastSortKey: "1,2,2"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: txs,
	}
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, &types.LastSortKey{Contract: "abc", SortKey: "1,2,1"})

	result := validator.validateLastSortKeys(block)

	require.Equal(t, "Rejected proposal: invalid last sort key", logger.Msg)
	require.False(t, result)
}

func TestValidateIndex(t *testing.T) {
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.validateIndex(0)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestValidateIndexNotFirst(t *testing.T) {
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.validateIndex(1)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: Arweave block must be in the first transaction in the sequencer block", logger.Msg)
}

func TestValidateArweaveBlockTx(t *testing.T) {
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block)
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.validateArweaveBlockTx(tx)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestValidateArweaveBlockTxNoMsgs(t *testing.T) {
	tx := test.CreateTxWithMsgs(t)
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.validateArweaveBlockTx(tx)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: transaction with Arweave block must have exactly one message", logger.Msg)
}

func TestValidateArweaveBlockTxToManyMsgs(t *testing.T) {
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block, &block)
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.validateArweaveBlockTx(tx)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: transaction with Arweave block must have exactly one message", logger.Msg)
}

func TestValidateArweaveBlockMsg(t *testing.T) {
	block := test.ArweaveBlock()
	validator, logger := validatorAndLogger(t, nil, block.BlockInfo, nil)

	result := validator.validateArweaveBlockMsg(&block)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestValidateArweaveBlockMsgWithoutHoursDelay(t *testing.T) {
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	block := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692357016,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}

	result := validator.validateArweaveBlockMsg(block)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: Arweave block should be one hour older than the sequencer block", logger.Msg)
}

func TestValidateArweaveBlockMsgWithoutNextHeight(t *testing.T) {
	block := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692353410,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	result := validator.validateArweaveBlockMsg(block)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: new height of the Arweave block is not the next value compared to the previous height", logger.Msg)
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
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: oldBlock,
	}, nil, nil)

	result := validator.validateArweaveBlockMsg(&types.MsgArweaveBlock{
		BlockInfo: newBlock,
	})

	require.False(t, result)
	require.Equal(t, "Rejected proposal: timestamp of the Arweave block is not later than the previous one", logger.Msg)
}

func TestValidateArweaveBlockMsgWithoutNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.validateArweaveBlockMsg(&block)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: the Validator did not fetch the Arweave block with given height", logger.Msg)
}

func TestValidateArweaveBlockMsgTimestampMismatchWithNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	nextBlockInfo := *block.BlockInfo
	nextBlockInfo.Timestamp += 1
	validator, logger := validatorAndLogger(t, nil, &nextBlockInfo, nil)

	result := validator.validateArweaveBlockMsg(&block)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: timestamp of the Arweave block does not match the timestamp of the block downloaded by the Validator", logger.Msg)
}

func TestValidateArweaveBlockMsgServerHashMismatchWithNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	nextBlockInfo := *block.BlockInfo
	nextBlockInfo.Hash = "abc"
	validator, logger := validatorAndLogger(t, nil, &nextBlockInfo, nil)

	result := validator.validateArweaveBlockMsg(&block)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: hash of the Arweave block does not match the hash of the block downloaded by the Validator", logger.Msg)
}

func TestArweaveBlockIsNotMissingGenesis(t *testing.T) {
	validator, logger := validatorAndLogger(t, nil, test.ArweaveBlock().BlockInfo, nil)
	validator.sequencerBlockHeader.Height = 0

	result := validator.checkArweaveBlockIsNotMissing(0)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestArweaveBlockIsNotMissingNotFirst(t *testing.T) {
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.checkArweaveBlockIsNotMissing(1)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestArweaveBlockIsNotMissingNotNext(t *testing.T) {
	block := test.ArweaveBlock()
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	result := validator.checkArweaveBlockIsNotMissing(0)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestArweaveBlockIsNotMissing(t *testing.T) {
	blockInfo := test.ArweaveBlock().BlockInfo
	blockInfo.Timestamp += 1692356017
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: blockInfo,
	}, blockInfo, nil)

	result := validator.checkArweaveBlockIsNotMissing(0)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestArweaveBlockIsMissing(t *testing.T) {
	blockInfo := test.ArweaveBlock().BlockInfo
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: blockInfo,
	}, blockInfo, nil)

	result := validator.checkArweaveBlockIsNotMissing(0)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: first transaction of the block should contain a transaction with the Arweave block", logger.Msg)
}

func TestCheckTransactionsLengthMismatch(t *testing.T) {
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", SortKey: "123"}}
	block := &types.MsgArweaveBlock{
		BlockInfo: test.ArweaveBlock().BlockInfo,
	}
	validator, logger := validatorAndLogger(t, nil, nil, nil)

	result := validator.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: incorrect number of transactions in the Arweave block", logger.Msg)
}

func TestCheckTransactionsIdMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "1234", SortKey: "1,2,3"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	result := validator.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: transaction id is not as expected", logger.Msg)
}

func TestCheckTransactionsContractMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abcd", Id: "123", SortKey: "1,2,3"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	result := validator.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: the contract of the transaction does not match the expected one", logger.Msg)
}

func TestCheckTransactionsSortKeyMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3,4"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	result := validator.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: transaction sort key is not as expected", logger.Msg)
}

func TestCheckTransactionsInvalidRandom(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3"}, LastSortKey: "1,1,1"}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	result := validator.checkTransactions(block, expectedTxs)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: transaction random value is not as expected", logger.Msg)
}

func TestCheckTransactions(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{
		Contract: "abc", Id: "123", SortKey: "1,2,3"}, LastSortKey: "1,1,1",
		Random: []byte{190, 96, 22, 62, 107, 198, 68, 216, 15, 189, 0, 227, 101, 238, 190, 27, 213, 120, 74, 38, 183, 173, 90, 197, 69, 66, 142, 157, 121, 160, 9, 117}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator, logger := validatorAndLogger(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	result := validator.checkTransactions(block, expectedTxs)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}
