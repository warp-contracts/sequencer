package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k *msgServer) ArweaveBlock(goCtx context.Context, msg *types.MsgArweaveBlock) (*types.MsgArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.setLastArweaveBlockInfo(ctx, msg)
	k.setContractPrevSortKeys(ctx, msg)

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

func (k *msgServer) setContractPrevSortKeys(ctx sdk.Context, msg *types.MsgArweaveBlock) {
	for _, tx := range msg.Transactions {
		prevSortKey := types.PrevSortKey{
			Contract: tx.Transaction.Contract,
			SortKey:  tx.Transaction.SortKey,
		}
		k.SetPrevSortKey(ctx, prevSortKey)
	}
}
