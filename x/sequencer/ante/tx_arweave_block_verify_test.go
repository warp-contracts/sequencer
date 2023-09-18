package ante

import (
	"github.com/stretchr/testify/require"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestArweaveBlockTxDecorator(t *testing.T) {
	ctx := sdk.Context{}.WithIsCheckTx(false)
	abdt := NewArweaveBlockTxDecorator()
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block)

	newCtx, err := abdt.AnteHandle(ctx, tx, false, nil)

	require.Equal(t, ctx, newCtx)
	require.NoError(t, err)
}

func TestArweaveBlockTxDecoratorNotByProposer(t *testing.T) {
	ctx := sdk.Context{}.WithIsCheckTx(true)
	abdt := NewArweaveBlockTxDecorator()
	block := test.ArweaveBlock()
	tx := test.CreateTxWithMsgs(t, &block)

	newCtx, err := abdt.AnteHandle(ctx, tx, false, nil)

	require.Equal(t, ctx, newCtx)
	require.ErrorIs(t, err, types.ErrArweaveBlockNotFromProposer)
}
