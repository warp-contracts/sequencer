package proposal

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func validatorAndMsg(t *testing.T) (*TxValidator, types.MsgDataItem) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	lastBlock := &types.LastArweaveBlock{
		ArweaveBlock:         arweaveBlockInfo,
		SequencerBlockHeight: 123,
	}
	validator := mockValidator(t, lastBlock, arweaveBlockInfo, nil)
	msg := test.ArweaveL2Interaction(t)
	return validator, msg
}

func TestCheckSortKey(t *testing.T) {
	validator, msg := validatorAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000000"

	err := validator.checkSortKey(&msg)

	require.NoError(t, err)
}

func TestCheckSortKeyNoSortKey(t *testing.T) {
	validator, msg := validatorAndMsg(t)

	err := validator.checkSortKey(&msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestCheckSortKeyInvalidArweaveBlock(t *testing.T) {
	validator, msg := validatorAndMsg(t)
	msg.SortKey = "000001431217,0000000000123,00000000"

	err := validator.checkSortKey(&msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestCheckSortKeyInvalidSequencerBlock(t *testing.T) {
	validator, msg := validatorAndMsg(t)
	msg.SortKey = "000001431216,0000000000124,00000000"

	err := validator.checkSortKey(&msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestCheckSortKeyInvalidIndex(t *testing.T) {
	validator, msg := validatorAndMsg(t)
	msg.SortKey = "000001431216,0000000000123,00000001"

	err := validator.checkSortKey(&msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestCheckSortKeyTwoMessagesInBlock(t *testing.T) {
	validator, msg := validatorAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	err := validator.checkSortKey(&msg)
	require.NoError(t, err)

	msg.SortKey = "000001431216,0000000000123,00000001"
	err = validator.checkSortKey(&msg)
	require.NoError(t, err)
}

func TestCheckSortKeyTwoSameSortKeysInBlock(t *testing.T) {
	validator, msg := validatorAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	err := validator.checkSortKey(&msg)
	require.NoError(t, err)

	msg.SortKey = "000001431216,0000000000123,00000000"
	err = validator.checkSortKey(&msg)
	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestCheckLastSortKeyEmptyKey(t *testing.T) {
	validator, msg := validatorAndMsg(t)

	err := validator.checkLastSortKey(&msg)

	require.NoError(t, err)
}

func TestCheckLastSortKeyNotEmptyFirstKey(t *testing.T) {
	validator, msg := validatorAndMsg(t)
	msg.LastSortKey = "1,2,3"

	err := validator.checkLastSortKey(&msg)

	require.ErrorIs(t, err, types.ErrInvalidLastSortKey)
}

func TestCheckLastSortKeyTwoMessagesInBlock(t *testing.T) {
	validator, msg := validatorAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	err := validator.checkLastSortKey(&msg)
	require.NoError(t, err)

	msg.SortKey = "000001431216,0000000000123,00000001"
	msg.LastSortKey = "000001431216,0000000000123,00000000"
	err = validator.checkLastSortKey(&msg)
	require.NoError(t, err)
}

func TestCheckLastSortKeyTwoMessagesInBlockInvalidKey(t *testing.T) {
	validator, msg := validatorAndMsg(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	err := validator.checkLastSortKey(&msg)
	require.NoError(t, err)

	msg.SortKey = "000001431216,0000000000123,00000001"
	msg.LastSortKey = "000001431216,0000000000123,00000001"
	err = validator.checkLastSortKey(&msg)
	require.ErrorIs(t, err, types.ErrInvalidLastSortKey)
}
