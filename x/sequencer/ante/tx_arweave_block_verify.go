package ante

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Checks if the transaction with the Arweave block is added by the Proposer.
type ArweaveBlockTxDecorator struct{}

func NewArweaveBlockTxDecorator() ArweaveBlockTxDecorator {
	return ArweaveBlockTxDecorator{}
}

func (abtd ArweaveBlockTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if isArweaveBlockTx(tx) {
		if ctx.IsCheckTx() {
			return ctx, errors.Wrapf(types.ErrArweaveBlockNotFromProposer,
				"transaction with arweave block can only be added by the Proposer")
		}

		// transaction with Arweave block does not require further validation within AnteHandler
		return ctx, nil
	}

	return next(ctx, tx, simulate)
}

func isArweaveBlockTx(tx sdk.Tx) bool {
	msgs := tx.GetMsgs()
	for _, msg := range msgs {
		_, isArweaveBlock := msg.(*types.MsgArweaveBlock)
		if isArweaveBlock {
			return true
		}
	}
	return false
}
