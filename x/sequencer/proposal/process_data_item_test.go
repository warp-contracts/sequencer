package proposal

import (
	"github.com/stretchr/testify/require"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func ctxHandlerLoggerAndMsg(t *testing.T) (sdk.Context, *processProposalHandler, *LoggerMock, types.MsgDataItem) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	lastBlock := &types.LastArweaveBlock{
		ArweaveBlock:         arweaveBlockInfo,
		SequencerBlockHeight: 123,
	}
	ctx, handler, logger := ctxHandlerAndLogger(t, lastBlock, arweaveBlockInfo)
	msg := test.ArweaveL2Interaction(t)
	return ctx, handler, logger, msg
}

func TestProcessProposalValidateDataItem(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000000"

	handler.initSortKeyForBlock(ctx)
	result := handler.processProposalValidateDataItem(ctx, &msg)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestProcessProposalValidateDataItemNoSortKey(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)
	result := handler.processProposalValidateDataItem(ctx, &msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestProcessProposalValidateDataItemInvalidArweaveBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431217,0000000000123,00000000"

	handler.initSortKeyForBlock(ctx)
	result := handler.processProposalValidateDataItem(ctx, &msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestProcessProposalValidateDataItemInvalidSequencerBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000124,00000000"

	handler.initSortKeyForBlock(ctx)
	result := handler.processProposalValidateDataItem(ctx, &msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestProcessProposalValidateDataItemInvalidIndex(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000001"

	handler.initSortKeyForBlock(ctx)
	result := handler.processProposalValidateDataItem(ctx, &msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestProcessProposalValidateDataItemTwoMessagesInBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := handler.processProposalValidateDataItem(ctx, &msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	msg.SortKey = "000001431216,0000000000123,00000001"
	result = handler.processProposalValidateDataItem(ctx, &msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestProcessProposalValidateDataItemTwoSameSortKeysInBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := handler.processProposalValidateDataItem(ctx, &msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	msg.SortKey = "000001431216,0000000000123,00000000"
	result = handler.processProposalValidateDataItem(ctx, &msg)
	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestProcessProposalValidateDataItemTwoBlocks(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	handler.initSortKeyForBlock(ctx)
	result := handler.processProposalValidateDataItem(ctx, &msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	header := ctx.BlockHeader()
	header.Height++
	ctx = ctx.WithBlockHeader(header)

	msg.SortKey = "000001431216,0000000000124,00000000"
	handler.initSortKeyForBlock(ctx)
	result = handler.processProposalValidateDataItem(ctx, &msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestProcessProposalValidateDataItemPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function should panic in case of uninitialized lastSortKey")
		}
	}()

	ctx, handler, _, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000000"

	handler.processProposalValidateDataItem(ctx, &msg)
}
