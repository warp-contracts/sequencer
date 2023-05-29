package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Checking if a transaction containing a DataItem has zero fees set.
func verifyFee(tx sdk.Tx) error {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "transaction is not of type FeeTx")
	}

	if feeTx.GetGas() != 0 {
		return sdkerrors.Wrapf(types.ErrNonZeroGas,
			"transaction with data item cannot have non-zero gas: %d", feeTx.GetGas())
	}

	for _, coin := range feeTx.GetFee() {
		if coin.Amount.Int64() != 0 {
			return sdkerrors.Wrapf(types.ErrNonZeroFee,
				"transaction with data item cannot have non-zero fee: %d", coin.String())
		}
	}

	if !feeTx.FeePayer().Empty() {
		return sdkerrors.Wrapf(types.ErrNotEmptyFeePayer,
			"transaction with data item cannot fee payer: %s", feeTx.FeePayer())
	}
	
	if !feeTx.FeeGranter().Empty() {
		return sdkerrors.Wrapf(types.ErrNotEmptyFeeGranter,
			"transaction with data item cannot fee granger: %s", feeTx.FeeGranter())
	}

	if err := verifyTip(tx); err != nil {
		return err
	}
	
	return nil
}

func verifyTip(sdkTx sdk.Tx) error {
	tipTx, ok := sdkTx.(tx.TipTx)
	if ok && tipTx.GetTip() != nil {
		return sdkerrors.Wrap(types.ErrNotEmptyTip, "transaction with data cannot have a tip")
	}
	return nil
}
