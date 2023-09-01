package ante

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// It checks if the transaction representing an Arweave block contains exacly one message: MsgArweaveBlock
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
	for _, msg := range msgs {
		_, isArweaveBlock := msg.(*types.MsgArweaveBlock)
		if isArweaveBlock {
			if len(msgs) > 1 {
				err := errors.Wrapf(types.ErrTooManyMessages,
					"transaction with arweave block can have only one message, and it has: %d", len(msgs))
				return err
			}
			return nil
		}
	}

	return nil

}
