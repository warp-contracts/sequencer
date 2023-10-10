package proposal

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func validatorLoggerAndMsg(t *testing.T) (*TxValidator, *LoggerMock, types.MsgDataItem) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	lastBlock := &types.LastArweaveBlock{
		ArweaveBlock:         arweaveBlockInfo,
		SequencerBlockHeight: 123,
	}
	validator, logger := validatorAndLogger(t, lastBlock, arweaveBlockInfo, nil)
	msg := test.ArweaveL2Interaction(t)
	return validator, logger, msg
}

func TestCheckSortKey(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000000"

	result := validator.checkSortKey(&msg)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestCheckSortKeyNoSortKey(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)

	result := validator.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: invalid sort key", logger.Msg)
}

func TestCheckSortKeyInvalidArweaveBlock(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)
	msg.SortKey = "000001431217,0000000000123,00000000"

	result := validator.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: invalid sort key", logger.Msg)
}

func TestCheckSortKeyInvalidSequencerBlock(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000124,00000000"

	result := validator.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: invalid sort key", logger.Msg)
}

func TestCheckSortKeyInvalidIndex(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000001"

	result := validator.checkSortKey(&msg)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: invalid sort key", logger.Msg)
}

func TestCheckSortKeyTwoMessagesInBlock(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := validator.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, "", logger.Msg)

	msg.SortKey = "000001431216,0000000000123,00000001"
	result = validator.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestCheckSortKeyTwoSameSortKeysInBlock(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := validator.checkSortKey(&msg)
	require.True(t, result)
	require.Equal(t, "", logger.Msg)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result = validator.checkSortKey(&msg)
	require.False(t, result)
	require.Equal(t, "Rejected proposal: invalid sort key", logger.Msg)
}

func TestCheckLastSortKeyEmptyKey(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)

	result := validator.checkLastSortKey(&msg)

	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestCheckLastSortKeyNotEmptyFirstKey(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)
	msg.LastSortKey = "1,2,3"

	result := validator.checkLastSortKey(&msg)

	require.False(t, result)
	require.Equal(t, "Rejected proposal: invalid last sort key", logger.Msg)
}

func TestCheckLastSortKeyTwoMessagesInBlock(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := validator.checkLastSortKey(&msg)
	require.True(t, result)
	require.Equal(t, "", logger.Msg)

	msg.SortKey = "000001431216,0000000000123,00000001"
	msg.LastSortKey = "000001431216,0000000000123,00000000"
	result = validator.checkLastSortKey(&msg)
	require.True(t, result)
	require.Equal(t, "", logger.Msg)
}

func TestCheckLastSortKeyTwoMessagesInBlockInvalidKey(t *testing.T) {
	validator, logger, msg := validatorLoggerAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	result := validator.checkLastSortKey(&msg)
	require.True(t, result)
	require.Equal(t, "", logger.Msg)

	msg.SortKey = "000001431216,0000000000123,00000001"
	msg.LastSortKey = "000001431216,0000000000123,00000001"
	result = validator.checkLastSortKey(&msg)
	require.False(t, result)
	require.Equal(t, "Rejected proposal: invalid last sort key", logger.Msg)
}
