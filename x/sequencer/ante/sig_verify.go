package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

type SigVerificationDecorator struct {
	standardSigVerificationDecorator ante.SigVerificationDecorator
}

func NewSigVerificationDecorator(standardSigVerificationDecorator ante.SigVerificationDecorator) SigVerificationDecorator {
	return SigVerificationDecorator{
		standardSigVerificationDecorator: standardSigVerificationDecorator,
	}
}

// The standard signature verification is only executed for transactions that do not have an Arweave DataItem
func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	dataItem, _ := GetDataItemMsg(tx)
	if dataItem != nil {
		return svd.AnteHandle(ctx, tx, simulate, next)
	}
	return next(ctx, tx, simulate)
}
