package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/test"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func lastArweaveBlock(ctx sdk.Context, k *keeper.Keeper, blockInfo *types.ArweaveBlockInfo) {
	lastArweaveBlock := types.LastArweaveBlock{
		ArweaveBlock:         blockInfo,
		SequencerBlockHeight: ctx.BlockHeight(),
	}
	k.SetLastArweaveBlock(ctx, lastArweaveBlock)
}

func TestDataItemMsgServer(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	k, ctx, srv := keeperCtxAndSrv(t, arweaveBlockInfo)
	lastArweaveBlock(ctx, k, arweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)
	msg := test.ArweaveL2Interaction(t)
	msg.SortKey = "000001431216,0000000000123,00000000"

	_, err := srv.DataItem(wctx, &msg)

	require.NoError(t, err)
}

func TestDataItemMsgServerNoSortKey(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	k, ctx, srv := keeperCtxAndSrv(t, arweaveBlockInfo)
	lastArweaveBlock(ctx, k, arweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)
	msg := test.ArweaveL2Interaction(t)

	_, err := srv.DataItem(wctx, &msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestDataItemMsgServerInvalidArweaveBlock(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	k, ctx, srv := keeperCtxAndSrv(t, arweaveBlockInfo)
	lastArweaveBlock(ctx, k, arweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)
	msg := test.ArweaveL2Interaction(t)
	msg.SortKey = "000001431217,0000000000123,00000000"

	_, err := srv.DataItem(wctx, &msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestDataItemMsgServerInvalidSequencerBlock(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	k, ctx, srv := keeperCtxAndSrv(t, arweaveBlockInfo)
	lastArweaveBlock(ctx, k, arweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)
	msg := test.ArweaveL2Interaction(t)
	msg.SortKey = "000001431216,0000000000124,00000000"

	_, err := srv.DataItem(wctx, &msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestDataItemMsgServerInvalidIndex(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	k, ctx, srv := keeperCtxAndSrv(t, arweaveBlockInfo)
	lastArweaveBlock(ctx, k, arweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)
	msg := test.ArweaveL2Interaction(t)
	msg.SortKey = "000001431216,0000000000123,00000001"

	_, err := srv.DataItem(wctx, &msg)

	require.ErrorIs(t, err, types.ErrInvalidSortKey)
}

func TestDataItemMsgServerTwoMessagesInBlock(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	k, ctx, srv := keeperCtxAndSrv(t, arweaveBlockInfo)
	lastArweaveBlock(ctx, k, arweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)
	msg := test.ArweaveL2Interaction(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	_, err := srv.DataItem(wctx, &msg)
	require.NoError(t, err)

	msg.SortKey = "000001431216,0000000000123,00000001"
	_, err = srv.DataItem(wctx, &msg)
	require.NoError(t, err)
}

func TestDataItemMsgServerTwoBlocks(t *testing.T) {
	arweaveBlockInfo := test.ArweaveBlock().BlockInfo
	k, ctx, srv := keeperCtxAndSrv(t, arweaveBlockInfo)
	lastArweaveBlock(ctx, k, arweaveBlockInfo)
	wctx := sdk.WrapSDKContext(ctx)
	msg := test.ArweaveL2Interaction(t)

	msg.SortKey = "000001431216,0000000000123,00000000"
	_, err := srv.DataItem(wctx, &msg)
	require.NoError(t, err)

	header := ctx.BlockHeader()
	header.Height++
	wctx = sdk.WrapSDKContext(ctx.WithBlockHeader(header))

	msg.SortKey = "000001431216,0000000000124,00000000"
	_, err = srv.DataItem(wctx, &msg)
	require.NoError(t, err)
}
