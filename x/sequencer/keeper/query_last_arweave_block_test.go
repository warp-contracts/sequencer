package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/testutil/nullify"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func TestLastArweaveBlockQuery(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	item := createTestLastArweaveBlock(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetLastArweaveBlockRequest
		response *types.QueryGetLastArweaveBlockResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLastArweaveBlockRequest{},
			response: &types.QueryGetLastArweaveBlockResponse{LastArweaveBlock: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LastArweaveBlock(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
