package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestLastArweaveBlockMsgServer(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgLastArweaveBlock{Creator: creator}
	_, err := srv.LastArweaveBlock(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetLastArweaveBlock(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}
