package app

import (
	"cosmossdk.io/errors"
	abci "github.com/cometbft/cometbft/abci/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/warp-contracts/sequencer/x/sequencer/ante"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (app *App) CheckTx(req abci.RequestCheckTx) abci.ResponseCheckTx {
	if req.Type == abci.CheckTxType_Recheck {
		ok, res := app.checkDataItemAlreadyInBlock(req.Tx)
		if !ok {
			return res
		}
	}
	res := app.BaseApp.CheckTx(req)
	app.logInvalidDataItemAfterRecheck(req, res)
	return res
}

func (app *App) checkDataItemAlreadyInBlock(txBytes []byte) (bool, abci.ResponseCheckTx) {
	tx, err := app.txConfig.TxDecoder()(txBytes)
	if err != nil {
		return false, sdkerrors.ResponseCheckTxWithEvents(err, 0, 0, nil, app.Trace())
	}
	dataItem, err := ante.GetL2Interaction(tx)
	if err != nil {
		return false, sdkerrors.ResponseCheckTxWithEvents(err, 0, 0, nil, app.Trace())
	}
	if dataItem != nil && app.BlockInteractions.Contains(dataItem) {
		err = errors.Wrap(types.ErrDataItemAlreadyInBlock,
			"The data item has already been added to the block")
		return false, sdkerrors.ResponseCheckTxWithEvents(err, 0, 0, nil, app.Trace())
	}
	return true, abci.ResponseCheckTx{}
}

func (app *App) logInvalidDataItemAfterRecheck(req abci.RequestCheckTx, res abci.ResponseCheckTx) {
	if req.Type != abci.CheckTxType_Recheck || res.Code == abci.CodeTypeOK {
		return
	}

	tx, err := app.txConfig.TxDecoder()(req.Tx)
	if err != nil {
		return
	}

	dataItem, err := ante.GetL2Interaction(tx)
	if dataItem == nil || err != nil {
		return
	}

	app.Logger().
		With("code", res.Code).
		With("log", res.Log).
		With("info", dataItem.GetInfo()).
		Info("Data item is no longer valid")
}
