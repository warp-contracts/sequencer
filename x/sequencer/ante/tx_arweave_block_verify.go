package ante

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// It checks if the transaction representing an Arweave block contains messages in the following order:
// - the first message is block information (MsgArweaveBlockInfo)
// - subsequent messages are L1 interactions (MsgDataItem)
type ArweaveBlockTxDecorator struct {
}

func NewArweaveBlockTxDecorator() ArweaveBlockTxDecorator {
	return ArweaveBlockTxDecorator{}
}

func (abtd ArweaveBlockTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	return ctx, verifyArweaveBlockTx(tx)
}

func verifyArweaveBlockTx(tx sdk.Tx) error {
	msgs := tx.GetMsgs()
	containsBlockInfoMsg := false

	for i, msg := range msgs {
		if containsBlockInfoMsg {
			if !isL1Interaction(msg) {
				return errors.Wrap(types.ErrInvalidArweaveBlockTx,
					"tx with an Arweave block can only contain block info and L1 interactions")
			}
		} else {
			if isArweaveBlockInfo(msg) {
				if i > 0 {
					return errors.Wrap(types.ErrInvalidArweaveBlockTx,
						"arweave block info must be the first message in the transaction")
				}
				containsBlockInfoMsg = true
			} else if isL1Interaction(msg) {
				return errors.Wrap(types.ErrInvalidArweaveBlockTx,
					"L1 interaction must be in tx after block info")
			}
		}
	}

	return nil

}

func isArweaveBlockInfo(msg sdk.Msg) bool {
	_, isBlockInfo := msg.(*types.MsgArweaveBlockInfo)
	return isBlockInfo
}

func isL1Interaction(msg sdk.Msg) bool {
	dataItem, isDataItem := msg.(*types.MsgDataItem)
	return isDataItem && dataItem.InteractionType == types.InteractionType_L1
}
