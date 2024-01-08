package keeper_test

import (
	"strconv"
	"testing"

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

func TestPrevSortKeyQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	msgs := createNPrevSortKey(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPrevSortKeyRequest
		response *types.QueryGetPrevSortKeyResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPrevSortKeyRequest{
				Contract: msgs[0].Contract,
			},
			response: &types.QueryGetPrevSortKeyResponse{PrevSortKey: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPrevSortKeyRequest{
				Contract: msgs[1].Contract,
			},
			response: &types.QueryGetPrevSortKeyResponse{PrevSortKey: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPrevSortKeyRequest{
				Contract: strconv.Itoa(100000),
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
			response, err := keeper.PrevSortKey(ctx, tc.request)
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

func TestPrevSortKeyQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	msgs := createNPrevSortKey(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPrevSortKeyRequest {
		return &types.QueryAllPrevSortKeyRequest{
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
			resp, err := keeper.PrevSortKeyAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PrevSortKey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PrevSortKey),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PrevSortKeyAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PrevSortKey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PrevSortKey),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PrevSortKeyAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PrevSortKey),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PrevSortKeyAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
