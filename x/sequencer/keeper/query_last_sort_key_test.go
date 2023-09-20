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

func TestLastSortKeyQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLastSortKey(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetLastSortKeyRequest
		response *types.QueryGetLastSortKeyResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLastSortKeyRequest{
				Contract: msgs[0].Contract,
			},
			response: &types.QueryGetLastSortKeyResponse{LastSortKey: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLastSortKeyRequest{
				Contract: msgs[1].Contract,
			},
			response: &types.QueryGetLastSortKeyResponse{LastSortKey: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLastSortKeyRequest{
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
			response, err := keeper.LastSortKey(wctx, tc.request)
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

func TestLastSortKeyQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLastSortKey(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLastSortKeyRequest {
		return &types.QueryAllLastSortKeyRequest{
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
			resp, err := keeper.LastSortKeyAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LastSortKey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LastSortKey),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LastSortKeyAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LastSortKey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LastSortKey),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LastSortKeyAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LastSortKey),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LastSortKeyAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
