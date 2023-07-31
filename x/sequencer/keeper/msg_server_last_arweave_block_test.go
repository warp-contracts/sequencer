package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestLastArweaveBlockMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.SequencerKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateLastArweaveBlock{Creator: creator}
	_, err := srv.CreateLastArweaveBlock(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetLastArweaveBlock(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestLastArweaveBlockMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateLastArweaveBlock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateLastArweaveBlock{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateLastArweaveBlock{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SequencerKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateLastArweaveBlock{Creator: creator}
			_, err := srv.CreateLastArweaveBlock(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateLastArweaveBlock(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetLastArweaveBlock(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestLastArweaveBlockMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteLastArweaveBlock
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteLastArweaveBlock{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteLastArweaveBlock{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SequencerKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateLastArweaveBlock(wctx, &types.MsgCreateLastArweaveBlock{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteLastArweaveBlock(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetLastArweaveBlock(ctx)
				require.False(t, found)
			}
		})
	}
}
