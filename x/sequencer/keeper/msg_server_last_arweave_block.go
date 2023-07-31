package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k msgServer) CreateLastArweaveBlock(goCtx context.Context, msg *types.MsgCreateLastArweaveBlock) (*types.MsgCreateLastArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetLastArweaveBlock(ctx)
	if isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var lastArweaveBlock = types.LastArweaveBlock{
		Creator: msg.Creator,
	}

	k.SetLastArweaveBlock(
		ctx,
		lastArweaveBlock,
	)
	return &types.MsgCreateLastArweaveBlockResponse{}, nil
}

func (k msgServer) UpdateLastArweaveBlock(goCtx context.Context, msg *types.MsgUpdateLastArweaveBlock) (*types.MsgUpdateLastArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLastArweaveBlock(ctx)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var lastArweaveBlock = types.LastArweaveBlock{
		Creator: msg.Creator,
	}

	k.SetLastArweaveBlock(ctx, lastArweaveBlock)

	return &types.MsgUpdateLastArweaveBlockResponse{}, nil
}

func (k msgServer) DeleteLastArweaveBlock(goCtx context.Context, msg *types.MsgDeleteLastArweaveBlock) (*types.MsgDeleteLastArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLastArweaveBlock(ctx)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLastArweaveBlock(ctx)

	return &types.MsgDeleteLastArweaveBlockResponse{}, nil
}
