package ante

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Validation of a transaction containing a L2 interaction.
// Such a transaction can have exactly one message, and all the values in this transaction are predetermined or derived from the DataItem.
// See: https://github.com/warp-contracts/sequencer/issues/8
type DataItemTxDecorator struct {
	ak *authkeeper.AccountKeeper
	bi *BlockInteractions
}

func NewDataItemTxDecorator(ak *authkeeper.AccountKeeper, bi *BlockInteractions) *DataItemTxDecorator {
	return &DataItemTxDecorator{ak, bi}
}

func (ditd *DataItemTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	dataItem, err := GetL2Interaction(tx)
	if err != nil {
		return ctx, err
	}

	if dataItem != nil {
		if err := ditd.verifyTxWithDataItem(ctx, tx, dataItem); err != nil {
			return ctx, err
		}
		// a valid transaction with a data item does not require further validation by the AnteHandler.
		return ctx, nil
	}

	return next(ctx, tx, simulate)
}

func GetL2Interaction(tx sdk.Tx) (*types.MsgDataItem, error) {
	msgs := tx.GetMsgs()

	for _, msg := range msgs {
		dataItem, isDataItem := msg.(*types.MsgDataItem)
		if isDataItem {
			if len(msgs) > 1 {
				err := errors.Wrapf(types.ErrInvalidMessagesNumber,
					"transaction with L2 interaction can have only one message, and it has: %d", len(msgs))
				return nil, err
			}
			return dataItem, nil
		}
	}

	return nil, nil
}

func isL2Interaction(tx sdk.Tx) bool {
	dataItem, err := GetL2Interaction(tx)
	return dataItem != nil && err == nil
}

func (ditd *DataItemTxDecorator) verifyTxWithDataItem(ctx sdk.Context, tx sdk.Tx, dataItem *types.MsgDataItem) (err error) {
	if ctx.IsReCheckTx() {
		if err = ditd.verifyAlreadyInBlock(ctx, dataItem); err != nil {
			return
		}
	} else {
		if err = verifyTxBody(tx); err != nil {
			return
		}

		if err = verifyFee(tx, dataItem); err != nil {
			return
		}

		if err = verifyContract(tx, dataItem); err != nil {
			return
		}

		if ctx.IsCheckTx() {
			// The data item is only validated before being inserted into the mempool 
			// and during the block validation in the `processProposalHandler`.
			if err = dataItem.Verify(); err != nil {
				return
			}
		}
	}

	if err := verifySignaturesAndNonce(ctx, ditd.ak, tx, dataItem); err != nil {
		return err
	}

	return nil
}

func (ditd *DataItemTxDecorator) verifyAlreadyInBlock(ctx sdk.Context, dataItem *types.MsgDataItem) (err error) {
	if ditd.bi.Contains(ctx.BlockHeight(), dataItem) {
		err = errors.Wrapf(types.ErrDataItemAlreadyInBlock, 
			"The data item has already been added to the block at height %d",
			ctx.BlockHeight(),
		)
	}
	return
}

func verifyContract(tx sdk.Tx, dataItem *types.MsgDataItem) error {
	_, err := dataItem.GetContractFromTags()
	return err
}
