package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/utils/arweave"
)

func (k *msgServer) ArweaveBlock(goCtx context.Context, msg *types.MsgArweaveBlock) (*types.MsgArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.setLastArweaveBlockInfo(ctx, msg)
	k.setContractLastSortKeys(ctx, msg)

	return &types.MsgArweaveBlockResponse{}, nil
}

func (k *msgServer) setLastArweaveBlockInfo(ctx sdk.Context, msg *types.MsgArweaveBlock) {
	newBlockInfo := *msg.BlockInfo

	lastArweaveBlock := types.LastArweaveBlock{
		ArweaveBlock:         &newBlockInfo,
		SequencerBlockHeight: ctx.BlockHeight(),
	}
	k.SetLastArweaveBlock(ctx, lastArweaveBlock)
}

func (k *msgServer) setContractLastSortKeys(ctx sdk.Context, msg *types.MsgArweaveBlock) {
	for _, tx := range msg.Transactions {
		lastSortKey := types.LastSortKey {
			Contract: arweave.Base64String(tx.Contract).Base64(),
			SortKey: tx.SortKey,
		}
		k.SetLastSortKey(ctx, lastSortKey)
	}
}
