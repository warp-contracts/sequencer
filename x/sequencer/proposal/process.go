package proposal

import (
	"time"

	"cosmossdk.io/log"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type processProposalHandler struct {
	txConfig       client.TxConfig
	logger         log.Logger
	blockValidator *BlockValidator
}

var (
	acceptResponse = &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}
	rejectResponse = &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}
)

func NewProcessProposalHandler(txConfig client.TxConfig, blockValidator *BlockValidator, logger log.Logger) sdk.ProcessProposalHandler {
	handler := &processProposalHandler{txConfig, logger.With("module", "block-validator"), blockValidator}
	return handler.process
}

// Validates the block proposal for the presence and correctness of transactions with the Arweave block,
// as well as the correctness of data items
func (h *processProposalHandler) process(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	now := time.Now()
	block := h.createBlock(ctx, req)
	if block == nil {
		return rejectResponse, nil
	}

	if err := h.blockValidator.ValidateBlock(block); err != nil {
		h.logger.
			With("err", err).
			With("height", req.Height).
			With("proposer", sdk.ConsAddress(req.ProposerAddress).String()).
			Info("Rejected proposal: invalid block")
		return rejectResponse, nil
	}

	h.logger.
		With("height", req.Height).
		With("number of txs", len(req.Txs)).
		With("proposer", sdk.ConsAddress(req.ProposerAddress).String()).
		With("execution time", time.Since(now).Milliseconds()).
		Info("Block accepted")
	return acceptResponse, nil
}

func (h *processProposalHandler) createBlock(ctx sdk.Context, req *abci.RequestProcessProposal) *Block {
	var txs []sdk.Tx
	for txIndex, txBytes := range req.Txs {
		tx, err := h.txConfig.TxDecoder()(txBytes)
		if err != nil {
			h.logger.Info("Rejected proposal: unable to decode the transaction", "err", err, "txIndex", txIndex)
			return nil
		}
		txs = append(txs, tx)
	}
	return &Block{ctx, txs}
}
