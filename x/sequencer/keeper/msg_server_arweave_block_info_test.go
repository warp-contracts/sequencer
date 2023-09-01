package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func keeperCtxAndSrv(t *testing.T) (*keeper.Keeper, sdk.Context, types.MsgServer) {
	k, ctx := keepertest.SequencerKeeper(t)
	blockHeader := ctx.BlockHeader()
	blockHeader.Time = time.Unix(1692357016, 0)
	srv := keeper.NewMsgServerImpl(*k)
	return k, ctx.WithBlockHeader(blockHeader), srv
}

func setNextArweaveBlock(k *keeper.Keeper, ctx sdk.Context, msgBlockInfo *types.MsgArweaveBlockInfo) {
	block := types.NextArweaveBlock {
		BlockInfo: &types.ArweaveBlockInfo{
			Height: msgBlockInfo.Height,
			Timestamp: msgBlockInfo.Timestamp,
			Hash: msgBlockInfo.Hash,
		},
	}
	k.SetNextArweaveBlock(ctx, block)
}

func TestArweaveBlockInfoMsgServer(t *testing.T) {
	k, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	expected := test.ArweaveBlockInfo()
	setNextArweaveBlock(k, ctx, &expected)

	_, err := srv.ArweaveBlockInfo(wctx, &expected)
	require.NoError(t, err)

	rst, found := k.GetLastArweaveBlock(ctx)
	require.True(t, found)
	require.Equal(t, expected.Height, rst.Height)
	require.Equal(t, expected.Timestamp, rst.Timestamp)
	require.Equal(t, expected.Hash, rst.Hash)
}

func TestArweaveBlockInfoMsgServerWithoutHoursDelay(t *testing.T) {
	_, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	lastArweaveBlock := &types.MsgArweaveBlockInfo{Creator: "creator", Height: 1431216, Timestamp: 1692357016, Hash: test.ExampleArweaveBlockHash}

	_, err := srv.ArweaveBlockInfo(wctx, lastArweaveBlock)
	require.ErrorIs(t, err, types.ErrArweaveBlockTimestampMismatch)
}

func TestArweaveBlockInfoMsgServerWithoutNextHeight(t *testing.T) {
	k, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	lastArweaveBlock := &types.MsgArweaveBlockInfo{Creator: "creator", Height: 1431216, Timestamp: 1692353410, Hash: test.ExampleArweaveBlockHash}
	setNextArweaveBlock(k, ctx, lastArweaveBlock)

	_, err := srv.ArweaveBlockInfo(wctx, lastArweaveBlock)
	require.NoError(t, err)

	lastArweaveBlock.Timestamp += 1
	_, err = srv.ArweaveBlockInfo(wctx, lastArweaveBlock)
	require.ErrorIs(t, err, types.ErrArweaveBlockHeightMismatch)
}

func TestArweaveBlockInfoMsgServerWithoutLaterTimestamp(t *testing.T) {
	k, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	lastArweaveBlock := &types.MsgArweaveBlockInfo{Creator: "creator", Height: 1431216, Timestamp: 1692353410, Hash: test.ExampleArweaveBlockHash}
	setNextArweaveBlock(k, ctx, lastArweaveBlock)

	_, err := srv.ArweaveBlockInfo(wctx, lastArweaveBlock)
	require.NoError(t, err)

	lastArweaveBlock.Height += 1
	_, err = srv.ArweaveBlockInfo(wctx, lastArweaveBlock)
	require.ErrorIs(t, err, types.ErrArweaveBlockTimestampMismatch)
}

func TestArweaveBlockInfoMsgServerWithoutNextArweaveBlock(t *testing.T) {
	_, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	arweaveBlockInfo := test.ArweaveBlockInfo()

	_, err := srv.ArweaveBlockInfo(wctx, &arweaveBlockInfo)
	require.ErrorIs(t, err, types.ErrInvalidArweaveBlockTx)
}

func TestArweaveBlockInfoMsgServerTimestampMismatchWithNextArweaveBlock(t *testing.T) {
	k, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	arweaveBlockInfo := test.ArweaveBlockInfo()
	nextArweaveBlockInfo := arweaveBlockInfo
	nextArweaveBlockInfo.Timestamp += 1
	setNextArweaveBlock(k, ctx, &nextArweaveBlockInfo)

	_, err := srv.ArweaveBlockInfo(wctx, &arweaveBlockInfo)
	require.ErrorIs(t, err, types.ErrArweaveBlockTimestampMismatch)
}

func TestArweaveBlockInfoMsgServerHashMismatchWithNextArweaveBlock(t *testing.T) {
	k, ctx, srv := keeperCtxAndSrv(t)
	wctx := sdk.WrapSDKContext(ctx)

	arweaveBlockInfo := test.ArweaveBlockInfo()
	nextArweaveBlockInfo := arweaveBlockInfo
	nextArweaveBlockInfo.Hash = []byte{1,2,3}
	setNextArweaveBlock(k, ctx, &nextArweaveBlockInfo)

	_, err := srv.ArweaveBlockInfo(wctx, &arweaveBlockInfo)
	require.ErrorIs(t, err, types.ErrArweaveBlockHashMismatch)
}
