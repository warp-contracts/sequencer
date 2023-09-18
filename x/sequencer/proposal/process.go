package proposal

import (
	"github.com/cometbft/cometbft/libs/log"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type processProposalHandler struct {
	keeper      *keeper.Keeper
	controller  controller.ArweaveBlocksController
	logger      log.Logger
	lastSortKey *types.SortKey
}

var (
	acceptResponse = abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}
	rejectResponse = abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}
)

// Validates the block proposal for the presence and correctness of transactions with the Arweave block,
// as well as the correctness of data items
func NewProcessProposalHandler(txConfig client.TxConfig, controller controller.ArweaveBlocksController, keeper *keeper.Keeper,
	logger log.Logger) sdk.ProcessProposalHandler {
	handler := &processProposalHandler{controller: controller, keeper: keeper, logger: logger}

	return func(ctx sdk.Context, req abci.RequestProcessProposal) abci.ResponseProcessProposal {
		handler.initSortKeyForBlock(ctx)
		for txIndex, txBytes := range req.Txs {
			tx, err := txConfig.TxDecoder()(txBytes)
			if err != nil {
				handler.rejectProposal("unable to decode the transaction", "err", err)
				return rejectResponse
			}
			if !handler.processProposalValidateTx(ctx, txIndex, tx) {
				return rejectResponse
			}
		}
		return acceptResponse
	}
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
