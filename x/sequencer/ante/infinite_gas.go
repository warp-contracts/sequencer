package ante

import(
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SetInfiniteGasMeterDecorator struct{}

func NewSetInfiniteGasMeterDecorator() SetInfiniteGasMeterDecorator {
	return SetInfiniteGasMeterDecorator{}
}

func (sigmd SetInfiniteGasMeterDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	return next(newCtx, tx, simulate)
}
