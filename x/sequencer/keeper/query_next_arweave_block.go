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

func (k Keeper) NextArweaveBlockAll(goCtx context.Context, req *types.QueryAllNextArweaveBlockRequest) (*types.QueryAllNextArweaveBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nextArweaveBlocks []types.NextArweaveBlock
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	nextArweaveBlockStore := prefix.NewStore(store, types.KeyPrefix(types.NextArweaveBlockKeyPrefix))

	pageRes, err := query.Paginate(nextArweaveBlockStore, req.Pagination, func(key []byte, value []byte) error {
		var nextArweaveBlock types.NextArweaveBlock
		if err := k.cdc.Unmarshal(value, &nextArweaveBlock); err != nil {
			return err
		}

		nextArweaveBlocks = append(nextArweaveBlocks, nextArweaveBlock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNextArweaveBlockResponse{NextArweaveBlock: nextArweaveBlocks, Pagination: pageRes}, nil
}

func (k Keeper) NextArweaveBlock(goCtx context.Context, req *types.QueryGetNextArweaveBlockRequest) (*types.QueryGetNextArweaveBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetNextArweaveBlock(
		ctx,
		req.Height,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNextArweaveBlockResponse{NextArweaveBlock: val}, nil
}
