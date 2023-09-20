package proposal

import (
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

func (h *processProposalHandler) processProposalValidateArweaveBlock(ctx sdk.Context, txIndex int, tx sdk.Tx, msg *types.MsgArweaveBlock) bool {
	accepted := h.validateIndex(txIndex) && h.validateArweaveBlockTx(tx) && h.validateArweaveBlockMsg(ctx, msg)
	if accepted {
		h.lastSortKey.IncreaseArweaveHeight()
	}
	return accepted
}

func (h *processProposalHandler) validateIndex(txIndex int) bool {
	if txIndex > 0 {
		return h.rejectProposal("Arweave block must be in the first transaction in the sequencer block",
			"transaction index", txIndex)
	}
	return true
}

func (h *processProposalHandler) validateArweaveBlockTx(tx sdk.Tx) bool {
	msgs := tx.GetMsgs()
	if len(msgs) != 1 {
		return h.rejectProposal("transaction with Arweave block must have exactly one message",
			"number of messages", len(msgs))
	}
	return true
}

func (h *processProposalHandler) validateArweaveBlockMsg(ctx sdk.Context, msg *types.MsgArweaveBlock) bool {
	if err := msg.ValidateBasic(); err != nil {
		return h.rejectProposal("invalid Arweave block message", "err", err)
	}

	newBlockInfo := &types.ArweaveBlockInfo{
		Height:    msg.BlockInfo.Height,
		Timestamp: msg.BlockInfo.Timestamp,
		Hash:      msg.BlockInfo.Hash,
	}

	return h.checkBlockIsOldEnough(ctx, newBlockInfo) &&
		h.compareBlockWithPreviousOne(ctx, newBlockInfo) &&
		h.compareWithNextBlock(ctx, msg)
}

func (h *processProposalHandler) checkBlockIsOldEnough(ctx sdk.Context, newBlockInfo *types.ArweaveBlockInfo) bool {
	arweaveBlockTimestamp := time.Unix(int64(newBlockInfo.Timestamp), 0)
	cosmosBlockTimestamp := ctx.BlockHeader().Time

	if !types.IsArweaveBlockOldEnough(ctx, newBlockInfo) {
		return h.rejectProposal("Arweave block should be one hour older than the sequencer block",
			"Arweave block timestamp", arweaveBlockTimestamp.UTC(),
			"sequencer block timestamp", cosmosBlockTimestamp.UTC())
	}
	return true
}

func (h *processProposalHandler) compareBlockWithPreviousOne(ctx sdk.Context, newValue *types.ArweaveBlockInfo) bool {
	oldValue, isFound := h.keeper.GetLastArweaveBlock(ctx)

	if !isFound {
		return true
	}

	if newValue.Height-oldValue.ArweaveBlock.Height != 1 {
		return h.rejectProposal("new height of the Arweave block is not the next value compared to the previous height")
	}

	if newValue.Timestamp <= oldValue.ArweaveBlock.Timestamp {
		return h.rejectProposal("timestamp of the Arweave block is not later than the previous one")
	}

	return true
}

func (h *processProposalHandler) compareWithNextBlock(ctx sdk.Context, block *types.MsgArweaveBlock) bool {
	nextArweaveBlock := h.controller.GetNextArweaveBlock(block.BlockInfo.Height)

	if nextArweaveBlock == nil {
		return h.rejectProposal("the Validator did not fetch the Arweave block with given height",
			"Arweave block height", block.BlockInfo.Height)
	}

	if block.BlockInfo.Timestamp != nextArweaveBlock.BlockInfo.Timestamp {
		return h.rejectProposal("timestamp of the Arweave block does not match the timestamp of the block downloaded by the Validator",
			"expected", nextArweaveBlock.BlockInfo.Timestamp, "actual", block.BlockInfo.Timestamp)
	}

	if block.BlockInfo.Hash != nextArweaveBlock.BlockInfo.Hash {
		return h.rejectProposal("hash of the Arweave block does not match the hash of the block downloaded by the Validator",
			"expected", string(nextArweaveBlock.BlockInfo.Hash), "actual", string(block.BlockInfo.Hash))
	}

	if transactionsDiffer(block.Transactions, nextArweaveBlock.Transactions) {
		return h.rejectProposal("Arweave block transactions do not match the block downloaded by the Validator transactions",
			"Arweave block height", block.BlockInfo.Height)
	}

	return true
}

func transactionsDiffer(transactions1 []*types.ArweaveTransaction, transactions2 []*types.ArweaveTransaction) bool {
	if len(transactions1) != len(transactions2) {
		return true
	}

	for i := 0; i < len(transactions1); i++ {
		tx1 := transactions1[i]
		tx2 := transactions2[i]
		if tx1.Id != tx2.Id || tx1.Contract != tx2.Contract {
			return true
		}
	}

	return false
}

func (h *processProposalHandler) checkArweaveBlockIsNotMissing(ctx sdk.Context, txIndex int) bool {
	if ctx.BlockHeader().Height == 0 || txIndex > 0 {
		return true
	}

	lastArweaveBlock := h.keeper.MustGetLastArweaveBlock(ctx)
	nextArweaveBlock := h.controller.GetNextArweaveBlock(lastArweaveBlock.ArweaveBlock.Height + 1)
	if nextArweaveBlock != nil && types.IsArweaveBlockOldEnough(ctx, nextArweaveBlock.BlockInfo) {
		return h.rejectProposal("first transaction of the block should contain a transaction with the Arweave block",
			"expected Arweave block height", nextArweaveBlock.BlockInfo.Height)
	}
	return true
}
