package proposal

import (
	"bytes"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func getArweaveBlockMsg(tx sdk.Tx) *types.MsgArweaveBlock {
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		arweaveBlock, isArweaveBlock := msg.(*types.MsgArweaveBlock)
		if isArweaveBlock {
			return arweaveBlock
		}
	}
	return nil
}

func (tv *TxValidator) validateSequentiallyArweaveBlock(txIndex int, tx sdk.Tx) bool {
	arweaveBlock := getArweaveBlockMsg(tx)
	if arweaveBlock != nil {
		tv.sortKey.IncreaseArweaveHeight()
		return tv.validateLastSortKeys(arweaveBlock)
	}
	return true
}

func (tv *TxValidator) validateLastSortKeys(block *types.MsgArweaveBlock) bool {
	for i, tx := range block.Transactions {
		expectedLastSortKey := tv.lastSortKeys.getAndStoreLastSortKey(tx.Transaction.Contract, tx.Transaction.SortKey)
		if tx.LastSortKey != expectedLastSortKey {
			return tv.rejectProposal("invalid last sort key",
				"Arweave block height", block.BlockInfo.Height, "transaction index", i, "expected", expectedLastSortKey, "actual", tx.LastSortKey)
		}
	}
	return true
}

func (tv *TxValidator) validateInParallelArweaveBlock(txIndex int, tx sdk.Tx) bool {
	arweaveBlock := getArweaveBlockMsg(tx)
	if arweaveBlock != nil {
		return tv.validateIndex(txIndex) && tv.validateArweaveBlockTx(tx) && tv.validateArweaveBlockMsg(arweaveBlock)
	} else {
		return tv.checkArweaveBlockIsNotMissing(txIndex)
	}
}

func (tv *TxValidator) validateIndex(txIndex int) bool {
	if txIndex > 0 {
		return tv.rejectProposal("Arweave block must be in the first transaction in the sequencer block",
			"transaction index", txIndex)
	}
	return true
}

func (tv *TxValidator) validateArweaveBlockTx(tx sdk.Tx) bool {
	msgs := tx.GetMsgs()
	if len(msgs) != 1 {
		return tv.rejectProposal("transaction with Arweave block must have exactly one message",
			"number of messages", len(msgs))
	}
	return true
}

func (tv *TxValidator) validateArweaveBlockMsg(msg *types.MsgArweaveBlock) bool {
	newBlockInfo := &types.ArweaveBlockInfo{
		Height:    msg.BlockInfo.Height,
		Timestamp: msg.BlockInfo.Timestamp,
		Hash:      msg.BlockInfo.Hash,
	}

	return tv.checkBlockIsOldEnough(newBlockInfo) &&
		tv.compareBlockWithPreviousOne(newBlockInfo) &&
		tv.compareWithNextBlock(msg)
}

func (tv *TxValidator) checkBlockIsOldEnough(newBlockInfo *types.ArweaveBlockInfo) bool {
	arweaveBlockTimestamp := time.Unix(int64(newBlockInfo.Timestamp), 0)
	sequencerBlockTimestamp := tv.sequencerBlockHeader.Time

	if !types.IsArweaveBlockOldEnough(tv.sequencerBlockHeader, newBlockInfo) {
		return tv.rejectProposal("Arweave block should be one hour older than the sequencer block",
			"Arweave block timestamp", arweaveBlockTimestamp.UTC(),
			"Sequencer block timestamp", sequencerBlockTimestamp.UTC())
	}
	return true
}

func (tv *TxValidator) compareBlockWithPreviousOne(newValue *types.ArweaveBlockInfo) bool {
	if newValue.Height-tv.lastArweaveBlock.ArweaveBlock.Height != 1 {
		return tv.rejectProposal("new height of the Arweave block is not the next value compared to the previous height")
	}

	if newValue.Timestamp <= tv.lastArweaveBlock.ArweaveBlock.Timestamp {
		return tv.rejectProposal("timestamp of the Arweave block is not later than the previous one")
	}

	return true
}

func (tv *TxValidator) compareWithNextBlock(block *types.MsgArweaveBlock) bool {
	if tv.nextArweaveBlock == nil {
		return tv.rejectProposal("the Validator did not fetch the Arweave block with given height",
			"Arweave block height", block.BlockInfo.Height)
	}

	if block.BlockInfo.Timestamp != tv.nextArweaveBlock.BlockInfo.Timestamp {
		return tv.rejectProposal("timestamp of the Arweave block does not match the timestamp of the block downloaded by the Validator",
			"expected", tv.nextArweaveBlock.BlockInfo.Timestamp, "actual", block.BlockInfo.Timestamp)
	}

	if block.BlockInfo.Hash != tv.nextArweaveBlock.BlockInfo.Hash {
		return tv.rejectProposal("hash of the Arweave block does not match the hash of the block downloaded by the Validator",
			"expected", string(tv.nextArweaveBlock.BlockInfo.Hash), "actual", string(block.BlockInfo.Hash))
	}

	return tv.checkTransactions(block, tv.nextArweaveBlock.Transactions)
}

func (tv *TxValidator) checkTransactions(block *types.MsgArweaveBlock, expectedTxs []*types.ArweaveTransaction) bool {
	if len(block.Transactions) != len(expectedTxs) {
		return tv.rejectProposal("incorrect number of transactions in the Arweave block",
			"Arweave block height", block.BlockInfo.Height, "expected", len(expectedTxs), "actual", len(block.Transactions))
	}

	for i := 0; i < len(expectedTxs); i++ {
		actualTx := block.Transactions[i]
		expectedTx := expectedTxs[i]

		if actualTx.Transaction.Id != expectedTx.Id {
			return tv.rejectProposal("transaction id is not as expected",
				"Arweave block height", block.BlockInfo.Height, "transaction index", i, "expected", expectedTx.Id, "actual", actualTx.Transaction.Id)
		}

		if actualTx.Transaction.Contract != expectedTx.Contract {
			return tv.rejectProposal("the contract of the transaction does not match the expected one",
				"Arweave block height", block.BlockInfo.Height, "transaction index", i, "expected", expectedTx.Contract, "actual", actualTx.Transaction.Contract)
		}

		if actualTx.Transaction.SortKey != expectedTx.SortKey {
			return tv.rejectProposal("transaction sort key is not as expected",
				"Arweave block height", block.BlockInfo.Height, "transaction index", i, "expected", expectedTx.SortKey, "actual", actualTx.Transaction.SortKey)
		}

		expectedRandom := generateRandomL1(actualTx.Transaction.SortKey)
		if !bytes.Equal(actualTx.Random, expectedRandom) {
			return tv.rejectProposal("transaction random value is not as expected",
				"Arweave block height", block.BlockInfo.Height, "transaction index", i, "expected", expectedRandom, "actual", actualTx.Random)
		}
	}

	return true
}

func (tv *TxValidator) checkArweaveBlockIsNotMissing(txIndex int) bool {
	if txIndex > 0 || tv.sequencerBlockHeader.Height == 0 {
		return true
	}

	if tv.nextArweaveBlock != nil && types.IsArweaveBlockOldEnough(tv.sequencerBlockHeader, tv.nextArweaveBlock.BlockInfo) {
		return tv.rejectProposal("first transaction of the block should contain a transaction with the Arweave block",
			"expected Arweave block height", tv.nextArweaveBlock.BlockInfo.Height)
	}
	return true
}
