package keeper

import (
	"bytes"
	"context"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

func (k *msgServer) ArweaveBlock(goCtx context.Context, msg *types.MsgArweaveBlock) (*types.MsgArweaveBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.setNewArweaveBlockInfo(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgArweaveBlockResponse{}, nil
}

func (k *msgServer) setNewArweaveBlockInfo(ctx sdk.Context, block *types.MsgArweaveBlock) error {
	var newBlockInfo = &types.ArweaveBlockInfo{
		Height:    block.BlockInfo.Height,
		Timestamp: block.BlockInfo.Timestamp,
		Hash:      block.BlockInfo.Hash,
	}

	if err := k.checkBlockIsOldEnough(ctx, newBlockInfo); err != nil {
		return err
	}

	if err := k.compareBlockWithPreviousOne(ctx, newBlockInfo); err != nil {
		return err
	}

	if err := k.compareWithNextBlock(ctx, block); err != nil {
		return err
	}

	k.setLastArweaveBlock(ctx, newBlockInfo)
	return nil
}

func (k *msgServer) checkBlockIsOldEnough(ctx sdk.Context, newBlockInfo *types.ArweaveBlockInfo) error {
	arweaveBlockTimestamp := time.Unix(int64(newBlockInfo.Timestamp), 0)
	cosmosBlockTimestamp := ctx.BlockHeader().Time

	if !types.IsArweaveBlockOldEnough(ctx, newBlockInfo) {
		return errors.Wrapf(types.ErrArweaveBlockTimestampMismatch,
			"The timestamp of the Arweave block (%s) should be one hour earlier than the Cosmos block (%s)",
			arweaveBlockTimestamp.UTC(), cosmosBlockTimestamp.UTC())
	}
	return nil
}

func (k *msgServer) compareBlockWithPreviousOne(ctx sdk.Context, newValue *types.ArweaveBlockInfo) error {
	oldValue, isFound := k.GetLastArweaveBlock(ctx)

	if !isFound {
		return nil
	}

	if newValue.Height-oldValue.ArweaveBlock.Height != 1 {
		return errors.Wrap(types.ErrArweaveBlockHeightMismatch,
			"The new height of the Arweave block is not the next value compared to the previous height")
	}

	if newValue.Timestamp <= oldValue.ArweaveBlock.Timestamp {
		return errors.Wrap(types.ErrArweaveBlockTimestampMismatch,
			"The timestamp of the Arweave block is not later than the previous one")
	}

	return nil
}

func (k *msgServer) compareWithNextBlock(ctx sdk.Context, block *types.MsgArweaveBlock) error {
	nextArweaveBlock := k.Controller.GetNextArweaveBlock(block.BlockInfo.Height)

	if nextArweaveBlock == nil {
		return errors.Wrapf(types.ErrInvalidArweaveBlockTx,
			"The validator did not fetch the Arweave block at height %d", block.BlockInfo.Height)
	}

	if block.BlockInfo.Timestamp != nextArweaveBlock.BlockInfo.Timestamp {
		return errors.Wrap(types.ErrArweaveBlockTimestampMismatch,
			"The timestamp of the Arweave block does not match the timestamp of the block downloaded by the Validator")
	}

	if !bytes.Equal(block.BlockInfo.Hash, nextArweaveBlock.BlockInfo.Hash) {
		return errors.Wrap(types.ErrArweaveBlockHashMismatch,
			"The hash of the Arweave block does not match the hash of the block downloaded by the Validator")
	}

	if transactionsDiffer(block.Transactions, nextArweaveBlock.Transactions) {
		return errors.Wrapf(types.ErrArweaveBlockTransactionsMismatch,
			"Arweave block transactions do not match the block downloaded by the Validator transactions for block %d",
			block.BlockInfo.Height)
	}

	return nil
}

func transactionsDiffer(transactions1 []*types.ArweaveTransaction, transactions2 []*types.ArweaveTransaction) bool {
	if len(transactions1) != len(transactions2) {
		return true
	}

	for i := 0; i < len(transactions1); i++ {
		tx1 := transactions1[i]
		tx2 := transactions2[i]
		if !bytes.Equal(tx1.Id, tx2.Id) || !bytes.Equal(tx1.Contract, tx2.Contract) {
			return true
		}
	}

	return false
}

func (k *msgServer) setLastArweaveBlock(ctx sdk.Context, newBlockInfo *types.ArweaveBlockInfo) {
	lastArweaveBlock := types.LastArweaveBlock {
		ArweaveBlock: newBlockInfo,
		SequencerBlockHeight: ctx.BlockHeight(),
	}
	k.SetLastArweaveBlock(ctx, lastArweaveBlock)
}
