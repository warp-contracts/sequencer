package proposal

import (
	"bytes"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func getDataItemMsg(tx sdk.Tx) *types.MsgDataItem {
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		dataItem, isDataItem := msg.(*types.MsgDataItem)
		if isDataItem {
			return dataItem
		}
	}
	return nil
}

func (tv *TxValidator) validateSequentiallyDataItem(txIndex int, tx sdk.Tx) error {
	dataItem := getDataItemMsg(tx)

	if dataItem != nil {
		if err := tv.checkSortKey(dataItem); err != nil {
			return err
		}
		return tv.checkLastSortKey(dataItem)
	}

	return nil
}

func (tv *TxValidator) checkSortKey(dataItem *types.MsgDataItem) error {
	expectedSortKey := tv.sortKey.GetNextValue()
	if expectedSortKey != dataItem.SortKey {
		return errors.Wrapf(types.ErrInvalidSortKey, "expected: %s, actual: %s", expectedSortKey, dataItem.SortKey)
	}

	return nil
}

func (tv *TxValidator) checkLastSortKey(dataItem *types.MsgDataItem) error {
	contract, err := dataItem.GetContractFromTags()
	if err != nil {
		return err
	}

	expectedLastSortKey := tv.lastSortKeys.getAndStoreLastSortKey(contract, dataItem.SortKey)
	if expectedLastSortKey != dataItem.LastSortKey {
		return errors.Wrapf(types.ErrInvalidLastSortKey, "expected: %s, actual: %s", expectedLastSortKey, dataItem.LastSortKey)
	}

	return nil
}

func (tv *TxValidator) validateInParallelDataItem(txIndex int, tx sdk.Tx) error {
	dataItem := getDataItemMsg(tx)

	if dataItem != nil {
		if err := dataItem.Verify(); err != nil {
			return err
		}
		return tv.checkRandom(dataItem)
	}

	return nil
}

func (tv *TxValidator) checkRandom(dataItem *types.MsgDataItem) error {
	expectedRandom := generateRandomL2(tv.sequencerBlockHeader.LastBlockId.Hash, dataItem.SortKey)
	if !bytes.Equal(dataItem.Random, expectedRandom) {
		return errors.Wrapf(types.ErrInvalidRandomValue, "expected: %s, actual: %s", expectedRandom, dataItem.Random)
	}

	return nil
}
