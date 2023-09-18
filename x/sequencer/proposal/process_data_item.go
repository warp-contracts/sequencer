package proposal

import (
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
	if err := msg.ValidateBasic(); err != nil {
		return h.rejectProposal("invalid data item message", "err", err)
	}

	if h.lastSortKey == nil || h.lastSortKey.SequencerHeight != ctx.BlockHeight() {
		h.lastSortKey = types.NewSortKey(h.keeper.MustGetLastArweaveBlock(ctx).ArweaveBlock.Height, ctx.BlockHeight())
	}

	expectedSortKey := h.lastSortKey.GetNextValue()
	if expectedSortKey != msg.SortKey {
		return h.rejectProposal("invalid sort key", "expected", expectedSortKey, "actual", msg.SortKey)
	}
	return true
}
