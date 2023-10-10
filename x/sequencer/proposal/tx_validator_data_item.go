package proposal

import (
	"bytes"

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

func (tv *TxValidator) validateSequentiallyDataItem(txIndex int, tx sdk.Tx) bool {
	dataItem := getDataItemMsg(tx)
	if dataItem != nil {
		return tv.checkSortKey(dataItem) && tv.checkLastSortKey(dataItem)
	}
	return true
}

func (tv *TxValidator) checkSortKey(msg *types.MsgDataItem) bool {
	expectedSortKey := tv.sortKey.GetNextValue()
	if expectedSortKey != msg.SortKey {
		return tv.rejectProposal("invalid sort key", "expected", expectedSortKey, "actual", msg.SortKey)
	}

	return true
}

func (tv *TxValidator) checkLastSortKey(msg *types.MsgDataItem) bool {
	contract, err := msg.GetContractFromTags()
	if err != nil {
		return tv.rejectProposal("invalid contract", "error", err)
	}
	expectedLastSortKey := tv.lastSortKeys.getAndStoreLastSortKey(contract, msg.SortKey)
	if expectedLastSortKey != msg.LastSortKey {
		return tv.rejectProposal("invalid last sort key", "expected", expectedLastSortKey, "actual", msg.LastSortKey)
	}

	return true
}

func (tv *TxValidator) validateInParallelDataItem(txIndex int, tx sdk.Tx) bool {
	dataItem := getDataItemMsg(tx)
	if dataItem != nil {
		return tv.checkDataItem(dataItem) && tv.checkRandom(dataItem)
	}
	return true
}

func (tv *TxValidator) checkDataItem(dataItem *types.MsgDataItem) bool {
	err := dataItem.Verify()
	if err != nil {
		return tv.rejectProposal("invalid data item message", "err", err)
	}
	return true
}

func (tv *TxValidator) checkRandom(dataItem *types.MsgDataItem) bool {
	expectedRandom := generateRandomL2(tv.sequencerBlockHeader.LastBlockId.Hash, dataItem.SortKey)
	if !bytes.Equal(dataItem.Random, expectedRandom) {
		return tv.rejectProposal("invalid random value", "expected", expectedRandom, "actual", dataItem.Random)
	}

	return true
}
