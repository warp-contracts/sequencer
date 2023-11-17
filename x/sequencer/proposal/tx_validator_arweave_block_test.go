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

type genesisLoaderMock struct{}

func (mock *genesisLoaderMock) LoadArweaveBlock() *types.GenesisArweaveBlock {
	return nil
}

func (mock *genesisLoaderMock) LoadPrevSortKeys() []types.PrevSortKey {
	return []types.PrevSortKey{}
}

func mockValidator(t *testing.T, lastBlock *types.LastArweaveBlock, nextBlock *types.ArweaveBlockInfo, prevSortKey *types.PrevSortKey) *TxValidator {
	keeper, ctx := keepertest.SequencerKeeper(t)

	if lastBlock == nil {
		keeper.SetLastArweaveBlock(ctx, types.LastArweaveBlock{ArweaveBlock: &types.ArweaveBlockInfo{Height: 1431215}})
	} else {
		keeper.SetLastArweaveBlock(ctx, *lastBlock)
	}

	if prevSortKey != nil {
		keeper.SetPrevSortKey(ctx, *prevSortKey)
	}

	blockHeader := ctx.BlockHeader()
	blockHeader.Time = time.Unix(1692357017, 0)
	blockHeader.Height = 123

	controller := controller.MockArweaveBlocksController(nextBlock)
	provider := NewArweaveBlockProvider(keeper, controller, &genesisLoaderMock{})

	return newTxValidator(ctx.WithBlockHeader(blockHeader), provider)
}

func TestValidatePrevSortKeysMismatch(t *testing.T) {
	txs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3"}, PrevSortKey: "1,2,2"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: txs,
	}
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, &types.PrevSortKey{Contract: "abc", SortKey: "1,2,1"})

	err := validator.validatePrevSortKeys(block)

	require.ErrorIs(t, err, types.ErrInvalidPrevSortKey)
}

func TestValidateIndex(t *testing.T) {
	validator := mockValidator(t, nil, nil, nil)

	err := validator.validateIndex(0)

	require.NoError(t, err)

}

func TestValidateIndexNotFirst(t *testing.T) {
	validator := mockValidator(t, nil, nil, nil)

	err := validator.validateIndex(1)

	require.ErrorIs(t, err, types.ErrInvalidTxIndex)
}

func TestValidateArweaveBlockTx(t *testing.T) {
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block)
	validator := mockValidator(t, nil, nil, nil)

	err := validator.validateArweaveBlockTx(tx)

	require.NoError(t, err)
}

func TestValidateArweaveBlockTxNoMsgs(t *testing.T) {
	tx := test.CreateTxWithMsgs(t)
	validator := mockValidator(t, nil, nil, nil)

	err := validator.validateArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrInvalidMessagesNumber)
}

func TestValidateArweaveBlockTxToManyMsgs(t *testing.T) {
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block, &block)
	validator := mockValidator(t, nil, nil, nil)

	err := validator.validateArweaveBlockTx(tx)

	require.ErrorIs(t, err, types.ErrInvalidMessagesNumber)
}

func TestValidateArweaveBlockMsg(t *testing.T) {
	block := test.ArweaveBlock()
	validator := mockValidator(t, nil, block.BlockInfo, nil)

	err := validator.validateArweaveBlockMsg(&block)

	require.NoError(t, err)
}

func TestValidateArweaveBlockMsgWithoutHoursDelay(t *testing.T) {
	validator := mockValidator(t, nil, nil, nil)

	block := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692357016,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}

	err := validator.validateArweaveBlockMsg(block)

	require.ErrorIs(t, err, types.ErrArweaveBlockNotOldEnough)
}

func TestValidateArweaveBlockMsgWithoutNextHeight(t *testing.T) {
	block := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692353410,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	err := validator.validateArweaveBlockMsg(block)

	require.ErrorIs(t, err, types.ErrBadArweaveHeight)
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
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: oldBlock,
	}, nil, nil)

	err := validator.validateArweaveBlockMsg(&types.MsgArweaveBlock{
		BlockInfo: newBlock,
	})

	require.ErrorIs(t, err, types.ErrBadArweaveTimestamp)
}

func TestValidateArweaveBlockMsgWithoutNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	validator := mockValidator(t, nil, nil, nil)

	err := validator.validateArweaveBlockMsg(&block)

	require.ErrorIs(t, err, types.ErrUnknownArweaveBlock)
}

func TestValidateArweaveBlockMsgTimestampMismatchWithNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	nextBlockInfo := *block.BlockInfo
	nextBlockInfo.Timestamp += 1
	validator := mockValidator(t, nil, &nextBlockInfo, nil)

	err := validator.validateArweaveBlockMsg(&block)

	require.ErrorIs(t, err, types.ErrBadArweaveTimestamp)
}

func TestValidateArweaveBlockMsgServerHashMismatchWithNextArweaveBlock(t *testing.T) {
	block := test.ArweaveBlock()
	nextBlockInfo := *block.BlockInfo
	nextBlockInfo.Hash = "abc"
	validator := mockValidator(t, nil, &nextBlockInfo, nil)

	err := validator.validateArweaveBlockMsg(&block)

	require.ErrorIs(t, err, types.ErrBadArweaveHash)
}

func TestArweaveBlockIsNotMissingGenesis(t *testing.T) {
	validator := mockValidator(t, nil, test.ArweaveBlock().BlockInfo, nil)
	validator.sequencerBlockHeader.Height = 0

	err := validator.checkArweaveBlockIsNotMissing()

	require.NoError(t, err)
}


func TestArweaveBlockIsNotMissingNoNext(t *testing.T) {
	block := test.ArweaveBlock()
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	err := validator.checkArweaveBlockIsNotMissing()

	require.NoError(t, err)
}

func TestArweaveBlockIsNotMissing(t *testing.T) {
	blockInfo := test.ArweaveBlock().BlockInfo
	blockInfo.Timestamp += 1692356017
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: blockInfo,
	}, blockInfo, nil)

	err := validator.checkArweaveBlockIsNotMissing()

	require.NoError(t, err)
}

func TestArweaveBlockIsMissing(t *testing.T) {
	blockInfo := test.ArweaveBlock().BlockInfo
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: blockInfo,
	}, blockInfo, nil)

	err := validator.checkArweaveBlockIsNotMissing()

	require.ErrorIs(t, err, types.ErrArweaveBlockMissing)
}

func TestCheckTransactionsLengthMismatch(t *testing.T) {
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", SortKey: "123"}}
	block := &types.MsgArweaveBlock{
		BlockInfo: test.ArweaveBlock().BlockInfo,
	}
	validator := mockValidator(t, nil, nil, nil)

	err := validator.checkTransactions(block, expectedTxs)

	require.ErrorIs(t, err, types.ErrInvalidTxNumber)
}

func TestCheckTransactionsIdMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "1234", SortKey: "1,2,3"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	err := validator.checkTransactions(block, expectedTxs)

	require.ErrorIs(t, err, types.ErrTxIdMismatch)
}

func TestCheckTransactionsContractMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abcd", Id: "123", SortKey: "1,2,3"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	err := validator.checkTransactions(block, expectedTxs)

	require.ErrorIs(t, err, types.ErrTxContractMismatch)
}

func TestCheckTransactionsSortKeyMismatch(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3,4"}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	err := validator.checkTransactions(block, expectedTxs)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestCheckTransactionsInvalidRandom(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{Contract: "abc", Id: "123", SortKey: "1,2,3"}, PrevSortKey: "1,1,1"}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	err := validator.checkTransactions(block, expectedTxs)

	require.ErrorIs(t, err, types.ErrInvalidRandomValue)
}

func TestCheckTransactions(t *testing.T) {
	actualTxs := []*types.ArweaveTransactionWithInfo{{Transaction: &types.ArweaveTransaction{
		Contract: "abc", Id: "123", SortKey: "1,2,3"}, PrevSortKey: "1,1,1",
		Random: []byte{190, 96, 22, 62, 107, 198, 68, 216, 15, 189, 0, 227, 101, 238, 190, 27, 213, 120, 74, 38, 183, 173, 90, 197, 69, 66, 142, 157, 121, 160, 9, 117}}}
	expectedTxs := []*types.ArweaveTransaction{{Contract: "abc", Id: "123", SortKey: "1,2,3"}}
	block := &types.MsgArweaveBlock{
		BlockInfo:    test.ArweaveBlock().BlockInfo,
		Transactions: actualTxs,
	}
	validator := mockValidator(t, &types.LastArweaveBlock{
		ArweaveBlock: block.BlockInfo,
	}, nil, nil)

	err := validator.checkTransactions(block, expectedTxs)

	require.NoError(t, err)
}
