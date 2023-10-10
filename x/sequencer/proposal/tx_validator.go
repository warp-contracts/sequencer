package proposal

import (
	"github.com/cometbft/cometbft/libs/log"
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
	logger               log.Logger
	sequencerBlockHeader tmproto.Header
	lastArweaveBlock     *types.LastArweaveBlock
	nextArweaveBlock     *types.NextArweaveBlock
	sortKey              *SortKey
	lastSortKeys         *LastSortKeys
}

func newTxValidator(ctx sdk.Context, keeper *keeper.Keeper, controller controller.ArweaveBlocksController, logger log.Logger) *TxValidator {
	lastArweaveBlock := keeper.MustGetLastArweaveBlock(ctx)
	nextArweaveBlock := controller.GetNextArweaveBlock(lastArweaveBlock.ArweaveBlock.Height + 1)
	sortKey := newSortKey(lastArweaveBlock.ArweaveBlock.Height, ctx.BlockHeight())
	lastSortKeys := newLastSortKeys(keeper, ctx)
	return &TxValidator{logger, ctx.BlockHeader(), &lastArweaveBlock, nextArweaveBlock, sortKey, lastSortKeys}
}

func (tv *TxValidator) validateSequentially(txIndex int, tx sdk.Tx) bool {
	return tv.validateSequentiallyArweaveBlock(txIndex, tx) && tv.validateSequentiallyDataItem(txIndex, tx)
}

func (tv *TxValidator) validateInParallel(txIndex int, tx sdk.Tx) bool {
	return tv.validateInParallelArweaveBlock(txIndex, tx) && tv.validateInParallelDataItem(txIndex, tx)
}

func (tv *TxValidator) rejectProposal(msg string, keyvals ...interface{}) bool {
	tv.logger.Info("Rejected proposal: "+msg, keyvals...)
	return false
}
