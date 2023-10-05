package proposal

import (
	"time"

	"github.com/cometbft/cometbft/libs/log"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
)

type processProposalHandler struct {
	txConfig     client.TxConfig
	keeper       *keeper.Keeper
	controller   controller.ArweaveBlocksController
	logger       log.Logger
	sortKey      *SortKey
	lastSortKeys *LastSortKeys
}

var (
	acceptResponse = abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}
	rejectResponse = abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}
)

func NewProcessProposalHandler(txConfig client.TxConfig, controller controller.ArweaveBlocksController, keeper *keeper.Keeper,
	logger log.Logger) sdk.ProcessProposalHandler {
	handler := &processProposalHandler{txConfig: txConfig, controller: controller, keeper: keeper, logger: logger}
	return handler.process
}

// Validates the block proposal for the presence and correctness of transactions with the Arweave block,
// as well as the correctness of data items
func (h *processProposalHandler) process(ctx sdk.Context, req abci.RequestProcessProposal) abci.ResponseProcessProposal {
	now := time.Now()
	h.initSortKeyForBlock(ctx)
	for txIndex, txBytes := range req.Txs {
		tx, err := h.txConfig.TxDecoder()(txBytes)
		if err != nil {
			h.rejectProposal("unable to decode the transaction", "err", err)
			return rejectResponse
		}
		if !h.processProposalValidateTx(ctx, txIndex, tx) {
			return rejectResponse
		}
	}

	ctx.Logger().
		With("height", req.Height).
		With("number of txs", len(req.Txs)).
		With("proposer", req.ProposerAddress).
		With("time", time.Since(now).Milliseconds()).
		Info("Block accepted")
	return acceptResponse
}

func (h *processProposalHandler) rejectProposal(msg string, keyvals ...interface{}) bool {
	h.logger.Info("Rejected proposal: "+msg, keyvals...)
	return false
}

func (h *processProposalHandler) processProposalValidateTx(ctx sdk.Context, txIndex int, tx sdk.Tx) bool {
	arweaveBlock := getArweaveBlockMsg(tx)
	if arweaveBlock != nil {
		return h.processProposalValidateArweaveBlock(ctx, txIndex, tx, arweaveBlock)
	}

	if !h.checkArweaveBlockIsNotMissing(ctx, txIndex) {
		return false
	}

	dataItem := getDataItemMsg(tx)
	if dataItem != nil {
		return h.processProposalValidateDataItem(ctx, dataItem)
	}

	return true
}
