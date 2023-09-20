package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LastSortKeyAll(goCtx context.Context, req *types.QueryAllLastSortKeyRequest) (*types.QueryAllLastSortKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lastSortKeys []types.LastSortKey
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	lastSortKeyStore := prefix.NewStore(store, types.KeyPrefix(types.LastSortKeyKeyPrefix))

	pageRes, err := query.Paginate(lastSortKeyStore, req.Pagination, func(key []byte, value []byte) error {
		var lastSortKey types.LastSortKey
		if err := k.cdc.Unmarshal(value, &lastSortKey); err != nil {
			return err
		}

		lastSortKeys = append(lastSortKeys, lastSortKey)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLastSortKeyResponse{LastSortKey: lastSortKeys, Pagination: pageRes}, nil
}

func (k Keeper) LastSortKey(goCtx context.Context, req *types.QueryGetLastSortKeyRequest) (*types.QueryGetLastSortKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetLastSortKey(
		ctx,
		req.Contract,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLastSortKeyResponse{LastSortKey: val}, nil
}
