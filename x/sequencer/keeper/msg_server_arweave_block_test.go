package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/ante"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func keeperCtxAndSrv(t *testing.T) (*keeper.Keeper, sdk.Context, types.MsgServer) {
	k, ctx := keepertest.SequencerKeeper(t)
	blockHeader := ctx.BlockHeader()
	blockHeader.Time = time.Unix(1692357017, 0)
	blockHeader.Height = 123

	bi := ante.NewBlockInteractions()
	bi.NewBlock()

	srv := keeper.NewMsgServerImpl(k, bi)
	return &k, ctx.WithBlockHeader(blockHeader), srv
}

func TestArweaveBlockMsgServer(t *testing.T) {
	expected := test.ArweaveBlock()
	k, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := srv.ArweaveBlock(wctx, &expected)
	require.NoError(t, err)

	result, found := k.GetLastArweaveBlock(ctx)
	require.True(t, found)
	require.Equal(t, expected.BlockInfo.Height, result.ArweaveBlock.Height)
	require.Equal(t, expected.BlockInfo.Timestamp, result.ArweaveBlock.Timestamp)
	require.Equal(t, expected.BlockInfo.Hash, result.ArweaveBlock.Hash)
	require.Equal(t, int64(123), result.SequencerBlockHeight)
	require.NotSame(t, expected.BlockInfo, result.ArweaveBlock)
}
