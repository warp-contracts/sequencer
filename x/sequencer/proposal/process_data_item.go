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

func (h *processProposalHandler) processProposalValidateDataItem(ctx sdk.Context, msg *types.MsgDataItem) bool {
	return h.checkSortKey(msg) && h.checkLastSortKey(msg) && h.checkRandom(ctx, msg)
}

func (h *processProposalHandler) checkSortKey(msg *types.MsgDataItem) bool {
	if h.sortKey == nil {
		panic("sortKey was not initialized")
	}

	expectedSortKey := h.sortKey.GetNextValue()
	if expectedSortKey != msg.SortKey {
		return h.rejectProposal("invalid sort key", "expected", expectedSortKey, "actual", msg.SortKey)
	}

	return true
}

func (h *processProposalHandler) checkLastSortKey(msg *types.MsgDataItem) bool {
	if h.lastSortKeys == nil {
		panic("lastSortKeys was not initialized")
	}

	contract, err := msg.GetContractFromTags()
	if err != nil {
		return h.rejectProposal("invalid contract", "error", err)
	}
	expectedLastSortKey := h.lastSortKeys.getAndStoreLastSortKey(contract, msg.SortKey)
	if expectedLastSortKey != msg.LastSortKey {
		return h.rejectProposal("invalid last sort key", "expected", expectedLastSortKey, "actual", msg.LastSortKey)
	}

	return true
}

func (h *processProposalHandler) initSortKeyForBlock(ctx sdk.Context) {
	h.sortKey = newSortKey(h.keeper.MustGetLastArweaveBlock(ctx).ArweaveBlock.Height, ctx.BlockHeight())
	h.lastSortKeys = newLastSortKeys(h.keeper, ctx)
}

func (h *processProposalHandler) checkRandom(ctx sdk.Context, msg *types.MsgDataItem) bool {
	expectedRandom := generateRandomL2(ctx.BlockHeader().LastBlockId.Hash, msg.SortKey)
	if !bytes.Equal(msg.Random, expectedRandom) {
		return h.rejectProposal("invalid random value", "expected", expectedRandom, "actual", msg.Random)
	}

	return true
}
