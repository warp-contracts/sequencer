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
	"github.com/warp-contracts/sequencer/x/sequencer/controller"
)

func keeperCtxAndSrv(t *testing.T, nextBlockInfo *types.ArweaveBlockInfo) (*keeper.Keeper, sdk.Context, types.MsgServer) {
	k, ctx := keepertest.SequencerKeeper(t)
	blockHeader := ctx.BlockHeader()
	blockHeader.Time = time.Unix(1692357017, 0)
	blockHeader.Height = 123
	
	c := controller.MockArweaveBlocksController(nextBlockInfo)
	
	srv := keeper.NewMsgServerImpl(*k, c)
	return k, ctx.WithBlockHeader(blockHeader), srv
}

func TestArweaveBlockMsgServer(t *testing.T) {
	expected := test.ArweaveBlock()
	k, ctx, srv := keeperCtxAndSrv(t, expected.BlockInfo)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := srv.ArweaveBlock(wctx, &expected)
	require.NoError(t, err)

	rst, found := k.GetLastArweaveBlock(ctx)
	require.True(t, found)
	require.Equal(t, expected.BlockInfo.Height, rst.ArweaveBlock.Height)
	require.Equal(t, expected.BlockInfo.Timestamp, rst.ArweaveBlock.Timestamp)
	require.Equal(t, expected.BlockInfo.Hash, rst.ArweaveBlock.Hash)
}

func TestArweaveBlockMsgServerWithoutHoursDelay(t *testing.T) {
	_, ctx, srv := keeperCtxAndSrv(t, nil)
	wctx := sdk.WrapSDKContext(ctx)

	arweaveBlock := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692357016,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}

	_, err := srv.ArweaveBlock(wctx, arweaveBlock)
	require.ErrorIs(t, err, types.ErrArweaveBlockTimestampMismatch)
}

func TestArweaveBlockMsgServerWithoutNextHeight(t *testing.T) {
	arweaveBlock := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692353410,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}
	_, ctx, srv := keeperCtxAndSrv(t, arweaveBlock.BlockInfo)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := srv.ArweaveBlock(wctx, arweaveBlock)
	require.NoError(t, err)

	arweaveBlock.BlockInfo.Timestamp += 1
	_, err = srv.ArweaveBlock(wctx, arweaveBlock)
	require.ErrorIs(t, err, types.ErrArweaveBlockHeightMismatch)
}

func TestArweaveBlockMsgServerWithoutLaterTimestamp(t *testing.T) {
	lastArweaveBlock := &types.MsgArweaveBlock{
		BlockInfo: &types.ArweaveBlockInfo{
			Height:    1431216,
			Timestamp: 1692353410,
			Hash:      test.ExampleArweaveBlockHash,
		},
	}
	_, ctx, srv := keeperCtxAndSrv(t, lastArweaveBlock.BlockInfo)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := srv.ArweaveBlock(wctx, lastArweaveBlock)
	require.NoError(t, err)

	lastArweaveBlock.BlockInfo.Height += 1
	_, err = srv.ArweaveBlock(wctx, lastArweaveBlock)
	require.ErrorIs(t, err, types.ErrArweaveBlockTimestampMismatch)
}

func TestArweaveBlockMsgServerWithoutNextArweaveBlock(t *testing.T) {
	_, ctx, srv := keeperCtxAndSrv(t, nil)
	wctx := sdk.WrapSDKContext(ctx)

	arweaveBlockInfo := test.ArweaveBlock()

	_, err := srv.ArweaveBlock(wctx, &arweaveBlockInfo)
	require.ErrorIs(t, err, types.ErrInvalidArweaveBlockTx)
}

func TestArweaveBlockMsgServerTimestampMismatchWithNextArweaveBlock(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock()
	nextArweaveBlockInfo := arweaveBlockInfo
	nextArweaveBlockInfo.BlockInfo.Timestamp += 1

	_, ctx, srv := keeperCtxAndSrv(t, nextArweaveBlockInfo.BlockInfo)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := srv.ArweaveBlock(wctx, &arweaveBlockInfo)
	require.ErrorIs(t, err, types.ErrArweaveBlockTimestampMismatch)
}

func TestArweaveBlockMsgServerHashMismatchWithNextArweaveBlock(t *testing.T) {

	arweaveBlockInfo := test.ArweaveBlock()
	nextArweaveBlockInfo := *arweaveBlockInfo.BlockInfo
	nextArweaveBlockInfo.Hash = []byte{1, 2, 3}

	_, ctx, srv := keeperCtxAndSrv(t, &nextArweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := srv.ArweaveBlock(wctx, &arweaveBlockInfo)
	require.ErrorIs(t, err, types.ErrArweaveBlockHashMismatch)
}
