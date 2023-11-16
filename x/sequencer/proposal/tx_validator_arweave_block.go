package proposal

import (
	"bytes"
	"time"

	"cosmossdk.io/errors"
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

func (tv *TxValidator) validateSequentiallyArweaveBlock(txIndex int, tx sdk.Tx) error {
	arweaveBlock := getArweaveBlockMsg(tx)
	if arweaveBlock != nil {
		tv.sortKey.IncreaseArweaveHeight()
		return tv.validatePrevSortKeys(arweaveBlock)
	}
	return nil
}

func (tv *TxValidator) validatePrevSortKeys(block *types.MsgArweaveBlock) error {
	for i, tx := range block.Transactions {
		expectedPrevSortKey := tv.prevSortKeys.getAndStorePrevSortKey(tx.Transaction.Contract, tx.Transaction.SortKey)
		if tx.PrevSortKey != expectedPrevSortKey {
			return errors.Wrapf(types.ErrInvalidPrevSortKey, "Arweave block height: %d, transaction index: %d, expected: %s, actual: %s",
				block.BlockInfo.Height, i, expectedPrevSortKey, tx.PrevSortKey)
		}
	}
	return nil
}

func (tv *TxValidator) validateInParallelArweaveBlock(txIndex int, tx sdk.Tx) error {
	arweaveBlock := getArweaveBlockMsg(tx)

	if arweaveBlock != nil {
		if err := tv.validateIndex(txIndex); err != nil {
			return err
		}

		if err := tv.validateArweaveBlockTx(tx); err != nil {
			return err
		}

		return tv.validateArweaveBlockMsg(arweaveBlock)
	}

	if txIndex == 0 {
		return tv.checkArweaveBlockIsNotMissing()
	}
	return nil
}

func (tv *TxValidator) validateIndex(txIndex int) error {
	if txIndex > 0 {
		return errors.Wrapf(types.ErrInvalidTxIndex, "Arweave block must be in the first transaction in the sequencer block, transaction index: %d", txIndex)
	}

	return nil
}

func (tv *TxValidator) validateArweaveBlockTx(tx sdk.Tx) error {
	numberOfMessages := len(tx.GetMsgs())
	if numberOfMessages != 1 {
		return errors.Wrapf(types.ErrInvalidMessagesNumber, "transaction with Arweave block must have exactly one message, number of messages: %d", numberOfMessages)
	}

	return nil
}

func (tv *TxValidator) validateArweaveBlockMsg(msg *types.MsgArweaveBlock) error {
	newBlockInfo := &types.ArweaveBlockInfo{
		Height:    msg.BlockInfo.Height,
		Timestamp: msg.BlockInfo.Timestamp,
		Hash:      msg.BlockInfo.Hash,
	}

	if err := tv.checkBlockIsOldEnough(newBlockInfo); err != nil {
		return err
	}

	if err := tv.compareBlockWithPreviousOne(newBlockInfo); err != nil {
		return err
	}

	return tv.compareWithNextBlock(msg)
}

func (tv *TxValidator) checkBlockIsOldEnough(newBlockInfo *types.ArweaveBlockInfo) error {
	if tv.sequencerBlockHeader.Height > 1 && !types.IsArweaveBlockOldEnough(tv.sequencerBlockHeader, newBlockInfo) {
		arweaveBlockTimestamp := time.Unix(int64(newBlockInfo.Timestamp), 0)
		sequencerBlockTimestamp := tv.sequencerBlockHeader.Time

		return errors.Wrapf(types.ErrArweaveBlockNotOldEnough,
			"Arweave block should be one hour older than the sequencer block, Arweave block timestamp: %s, sequencer block timestamp: %s",
			arweaveBlockTimestamp.UTC().Format(time.DateTime), sequencerBlockTimestamp.UTC().Format(time.DateTime))
	}
	return nil
}

func (tv *TxValidator) compareBlockWithPreviousOne(newValue *types.ArweaveBlockInfo) error {
	if newValue.Height-tv.lastArweaveBlock.ArweaveBlock.Height != 1 {
		return errors.Wrapf(types.ErrBadArweaveHeight,
			"new height (%d) of the Arweave block is not the next value compared to the previous height (%d)",
			newValue.Height, tv.lastArweaveBlock.ArweaveBlock.Height)
	}

	if newValue.Timestamp <= tv.lastArweaveBlock.ArweaveBlock.Timestamp {
		return errors.Wrapf(types.ErrBadArweaveTimestamp,
			"timestamp of the Arweave block (%d) is not later than the previous one (%d)",
			newValue.Timestamp, tv.lastArweaveBlock.ArweaveBlock.Timestamp)
	}

	return nil
}

func (tv *TxValidator) compareWithNextBlock(block *types.MsgArweaveBlock) error {
	if tv.nextArweaveBlock == nil {
		return errors.Wrapf(types.ErrUnknownArweaveBlock, "the Validator did not fetch the Arweave block with height: %d", block.BlockInfo.Height)
	}

	if block.BlockInfo.Timestamp != tv.nextArweaveBlock.BlockInfo.Timestamp {
		return errors.Wrapf(types.ErrBadArweaveTimestamp,
			"timestamp of the Arweave block does not match the timestamp of the block downloaded by the Validator, expected: %d, actual: %d",
			tv.nextArweaveBlock.BlockInfo.Timestamp, block.BlockInfo.Timestamp)
	}

	if block.BlockInfo.Hash != tv.nextArweaveBlock.BlockInfo.Hash {
		return errors.Wrapf(types.ErrBadArweaveHash,
			"hash of the Arweave block does not match the hash of the block downloaded by the Validator, expected: %d, actual: %d",
			tv.nextArweaveBlock.BlockInfo.Timestamp, block.BlockInfo.Timestamp)
	}

	return tv.checkTransactions(block, tv.nextArweaveBlock.Transactions)
}

func (tv *TxValidator) checkTransactions(block *types.MsgArweaveBlock, expectedTxs []*types.ArweaveTransaction) error {
	if len(block.Transactions) != len(expectedTxs) {
		return errors.Wrapf(types.ErrInvalidTxNumber,
			"incorrect number of transactions in the Arweave block with height: %d, expected: %d, actual: %d",
			block.BlockInfo.Height, len(expectedTxs), len(block.Transactions))
	}

	for i := 0; i < len(expectedTxs); i++ {
		actualTx := block.Transactions[i]
		expectedTx := expectedTxs[i]

		if actualTx.Transaction.Id != expectedTx.Id {
			return errors.Wrapf(types.ErrTxIdMismatch,
				"transaction id (%s) is not as expected (%s), Arweave block height: %d, transaction index: %d",
				actualTx.Transaction.Id, expectedTx.Id, block.BlockInfo.Height, i)
		}

		if actualTx.Transaction.Contract != expectedTx.Contract {
			return errors.Wrapf(types.ErrTxContractMismatch,
				"the contract of the transaction (%s) does not match the expected one (%s), Arweave block height: %d, transaction index: %d",
				actualTx.Transaction.Contract, expectedTx.Contract, block.BlockInfo.Height, i)
		}

		if actualTx.Transaction.SortKey != expectedTx.SortKey {
			return errors.Wrapf(types.ErrInvalidSortKey,
				"transaction sort key (%s) is not as expected (%s), Arweave block height: %d, transaction index: %d",
				actualTx.Transaction.SortKey, expectedTx.SortKey, block.BlockInfo.Height, i)
		}

		expectedRandom := generateRandomL1(actualTx.Transaction.SortKey)
		if !bytes.Equal(actualTx.Random, expectedRandom) {
			return errors.Wrapf(types.ErrInvalidRandomValue,
				"transaction random value (%s) is not as expected (%s), Arweave block height: %d, transaction index: %d",
				actualTx.Random, expectedRandom, block.BlockInfo.Height, i)
		}
	}

	return nil
}

func (tv *TxValidator) checkArweaveBlockIsNotMissing() error {
	if tv.sequencerBlockHeader.Height == 0 {
		return nil
	}

	if tv.nextArweaveBlock != nil && types.IsArweaveBlockOldEnough(tv.sequencerBlockHeader, tv.nextArweaveBlock.BlockInfo) {
		return errors.Wrapf(types.ErrArweaveBlockMissing,
			"first transaction of the block should contain a transaction with the Arweave block with height: %d",
			tv.nextArweaveBlock.BlockInfo.Height)
	}
	return nil
}
