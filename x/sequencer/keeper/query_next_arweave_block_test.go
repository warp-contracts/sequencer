package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/warp-contracts/sequencer/testutil/keeper"
	"github.com/warp-contracts/sequencer/testutil/nullify"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestNextArweaveBlockQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNextArweaveBlock(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetNextArweaveBlockRequest
		response *types.QueryGetNextArweaveBlockResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNextArweaveBlockRequest{
				Height: msgs[0].GetHeightString(),
			},
			response: &types.QueryGetNextArweaveBlockResponse{NextArweaveBlock: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNextArweaveBlockRequest{
				Height: msgs[1].GetHeightString(),
			},
			response: &types.QueryGetNextArweaveBlockResponse{NextArweaveBlock: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNextArweaveBlockRequest{
				Height: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NextArweaveBlock(wctx, tc.request)
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

func TestNextArweaveBlockQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNextArweaveBlock(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllNextArweaveBlockRequest {
		return &types.QueryAllNextArweaveBlockRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NextArweaveBlockAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NextArweaveBlock), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NextArweaveBlock),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NextArweaveBlockAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NextArweaveBlock), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NextArweaveBlock),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NextArweaveBlockAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NextArweaveBlock),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NextArweaveBlockAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
