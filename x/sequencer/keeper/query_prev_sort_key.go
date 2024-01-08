package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PrevSortKeyAll(ctx context.Context, req *types.QueryAllPrevSortKeyRequest) (*types.QueryAllPrevSortKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var prevSortKeys []types.PrevSortKey

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	prevSortKeyStore := prefix.NewStore(store, types.KeyPrefix(types.PrevSortKeyKeyPrefix))

	pageRes, err := query.Paginate(prevSortKeyStore, req.Pagination, func(key []byte, value []byte) error {
		var prevSortKey types.PrevSortKey
		if err := k.cdc.Unmarshal(value, &prevSortKey); err != nil {
			return err
		}

		prevSortKeys = append(prevSortKeys, prevSortKey)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPrevSortKeyResponse{PrevSortKey: prevSortKeys, Pagination: pageRes}, nil
}

func (k Keeper) PrevSortKey(ctx context.Context, req *types.QueryGetPrevSortKeyRequest) (*types.QueryGetPrevSortKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetPrevSortKey(
		ctx,
		req.Contract,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPrevSortKeyResponse{PrevSortKey: val}, nil
}
