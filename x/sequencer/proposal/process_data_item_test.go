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

func TestCheckSortKey(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000000"

	handler.initSortKeyForBlock(ctx)
	result := handler.checkSortKey(&msg)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestCheckSortKeyNoSortKey(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)
	result := handler.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestCheckSortKeyInvalidArweaveBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431217,0000000000123,00000000"

	handler.initSortKeyForBlock(ctx)
	result := handler.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestCheckSortKeyInvalidSequencerBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000124,00000000"

	handler.initSortKeyForBlock(ctx)
	result := handler.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestCheckSortKeyInvalidIndex(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000001"

	handler.initSortKeyForBlock(ctx)
	result := handler.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestCheckSortKeyTwoMessagesInBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := handler.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	msg.SortKey = "000001431216,0000000000123,00000001"
	result = handler.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestCheckSortKeyTwoSameSortKeysInBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := handler.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	msg.SortKey = "000001431216,0000000000123,00000000"
	result = handler.checkSortKey(&msg)
	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid sort key")
}

func TestCheckSortKeyTwoBlocks(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	handler.initSortKeyForBlock(ctx)
	result := handler.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	header := ctx.BlockHeader()
	header.Height++
	ctx = ctx.WithBlockHeader(header)

	msg.SortKey = "000001431216,0000000000124,00000000"
	handler.initSortKeyForBlock(ctx)
	result = handler.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestCheckSortKeyPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function should panic in case of uninitialized sortKey")
		}
	}()

	_, handler, _, msg := ctxHandlerLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000000"

	handler.checkSortKey(&msg)
}

func TestCheckLastSortKeyEmptyKey(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)
	result := handler.checkLastSortKey(&msg)

	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestCheckLastSortKeyNotEmptyFirstKey(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)
	msg.LastSortKey = "1,2,3"

	handler.initSortKeyForBlock(ctx)
	result := handler.checkLastSortKey(&msg)

	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid last sort key")
}

func TestCheckLastSortKeyTwoMessagesInBlock(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := handler.checkLastSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	msg.SortKey = "000001431216,0000000000123,00000001"
	msg.LastSortKey = "000001431216,0000000000123,00000000"
	result = handler.checkLastSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")
}

func TestCheckLastSortKeyTwoMessagesInBlockInvalidKey(t *testing.T) {
	ctx, handler, logger, msg := ctxHandlerLoggerAndMsg(t)

	handler.initSortKeyForBlock(ctx)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := handler.checkLastSortKey(&msg)
	require.True(t, result)
	require.Equal(t, logger.Msg, "")

	msg.SortKey = "000001431216,0000000000123,00000001"
	msg.LastSortKey = "000001431216,0000000000123,00000001"
	result = handler.checkLastSortKey(&msg)
	require.False(t, result)
	require.Equal(t, logger.Msg, "Rejected proposal: invalid last sort key")
}

func TestCheckLastSortKeyPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function should panic in case of uninitialized lastSortKey")
		}
	}()

	_, handler, _, msg := ctxHandlerLoggerAndMsg(t)
	handler.checkLastSortKey(&msg)
}
