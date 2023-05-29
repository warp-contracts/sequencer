package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Validation of a transaction containing a DataItem. 
// Such a transaction can have exactly one message, and all the values in this transaction are predetermined or derived from the DataItem.
// See: https://github.com/warp-contracts/sequencer/issues/8
type DataItemTxDecorator struct {
	ak sdkante.AccountKeeper
}

func NewDataItemTxDecorator(ak sdkante.AccountKeeper) DataItemTxDecorator {
	return DataItemTxDecorator{
		ak: ak,
	}
}

func (ditd DataItemTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	dataItem, err := GetDataItemMsg(tx)
	if err != nil {
		return ctx, err
	}

	if dataItem != nil {
		if err := verifyTxWithDataItem(ctx, ditd.ak, tx, dataItem); err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

func GetDataItemMsg(tx sdk.Tx) (*types.MsgDataItem, error) {
	msgs := tx.GetMsgs()

	if len(msgs) > 0 {
		dataItem, isDataItem := msgs[0].(*types.MsgDataItem)
		if isDataItem {
			if len(msgs) > 1 {
				err := sdkerrors.Wrapf(types.ErrToManyDataItems,
					"transaction with data item can have only one message, and it has: %d", len(msgs))
				return nil, err
			}
			return dataItem, nil
		}
	}

	return nil, nil
}

func verifyTxWithDataItem(ctx sdk.Context, ak sdkante.AccountKeeper, tx sdk.Tx, dataItem *types.MsgDataItem) error {
	if err := verifyTxBody(tx); err != nil {
		return err
	}

	if err := verifySignatures(ctx, ak, tx, dataItem); err != nil {
		return err
	}

	if err := verifyFee(tx); err != nil {
		return err
	}

	return nil
}
