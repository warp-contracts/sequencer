package proposal

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// It validates the transaction and provides two methods for validation: 
// - sequential - in the order of transactions in the block
// - parallel - independently of the validation of other transactions
type TxValidator struct {
	sequencerBlockHeader tmproto.Header
	lastArweaveBlock     *types.LastArweaveBlock
	nextArweaveBlock     *types.NextArweaveBlock
	sortKey              *SortKey
	prevSortKeys         *PrevSortKeys
}


func newTxValidator(ctx sdk.Context, keeper *keeper.Keeper, controller controller.ArweaveBlocksController) *TxValidator {
	lastArweaveBlock := keeper.MustGetLastArweaveBlock(ctx)
	nextArweaveBlock := controller.GetNextArweaveBlock(lastArweaveBlock.ArweaveBlock.Height + 1)
	sortKey := newSortKey(lastArweaveBlock.ArweaveBlock.Height, ctx.BlockHeight())
	prevSortKeys := newPrevSortKeys(keeper, ctx)
	return &TxValidator{ctx.BlockHeader(), &lastArweaveBlock, nextArweaveBlock, sortKey, prevSortKeys}
}

func (tv *TxValidator) validateSequentially(txIndex int, tx sdk.Tx) *InvalidTxError {
	if err := tv.validateSequentiallyArweaveBlock(txIndex, tx); err != nil {
		return InvalidArweaveError(err)
	}
	if err := tv.validateSequentiallyDataItem(txIndex, tx); err != nil {
		return InvalidDataItemError(err)
	}
	return nil
}

func (tv *TxValidator) validateInParallel(txIndex int, tx sdk.Tx) *InvalidTxError {
	if err := tv.validateInParallelArweaveBlock(txIndex, tx); err != nil {
		return InvalidArweaveError(err)
	}
	if err := tv.validateInParallelDataItem(txIndex, tx); err != nil {
		return InvalidDataItemError(err)
	}
	return nil
}
