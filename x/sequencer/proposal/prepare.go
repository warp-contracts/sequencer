package proposal

import (
	"fmt"
	math_bits "math/bits"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/ante"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type prepareProposalHandler struct {
	provider *ArweaveBlockProvider
	txConfig client.TxConfig
}

func NewPrepareProposalHandler(provider *ArweaveBlockProvider, txConfig client.TxConfig) sdk.PrepareProposalHandler {
	handler := &prepareProposalHandler{provider, txConfig}
	return handler.prepare
}

// Sets sort keys and random value for all L2 interactions and adds an Arweave block transaction to the beginning of the block if needed.
func (h *prepareProposalHandler) prepare(ctx sdk.Context, req abci.RequestPrepareProposal) abci.ResponsePrepareProposal {
	// For logging
	now := time.Now()

	// Helpert struct for assigning prev sort keys
	prevSortKeys := newPrevSortKeys(h.provider.keeper, ctx)

	var (
		// How much space do transactions occupy, there's a limit
		size int64

		// Helper structure for generating sort keys
		sortKey *SortKey

		// Encoded tx that will end up in the block
		result [][]byte
	)

	firstBlock := req.Height == 1
	// Get the last block that is stored in the blockchain
	lastBlock := h.provider.getLastArweaveBlock(ctx, firstBlock)
	// Try getting the next Arweave block, it may not be there or it may be too fresh to use
	nextArweaveBlock := h.provider.getNextArweaveBlock(lastBlock.ArweaveBlock.Height+1, firstBlock)

	if nextArweaveBlock != nil &&
		(firstBlock || types.IsArweaveBlockOldEnough(ctx.BlockHeader(), nextArweaveBlock.BlockInfo)) {
		// Sort keys are generated for the new block
		sortKey = newSortKey(lastBlock.ArweaveBlock.Height+1, ctx.BlockHeight())
		result = make([][]byte, 0, len(req.Txs)+1)

		// There's a new Arweave block, add it as the first tx in sequencer's block
		arweaveBlockTx := h.createArweaveTx(ctx, nextArweaveBlock, prevSortKeys)
		result = append(result, arweaveBlockTx)
		size += protoTxSize(arweaveBlockTx)
		if size > req.MaxTxBytes {
			panic(fmt.Errorf("MaxTxBytes limit (%d) is too small! It is smaller than the size of the Arweave block (%d)",
				req.MaxTxBytes, size))
		}
	} else {
		// Sort keys are generated for the last arweave block that is already in the blockchain
		sortKey = newSortKey(lastBlock.ArweaveBlock.Height, ctx.BlockHeight())
		result = make([][]byte, 0, len(req.Txs))
	}

	// Add transactions that waited in the mempool
	for i := range req.Txs {
		// Fill in helper data
		txBytes := h.prepareDataItem(ctx.BlockHeader().LastBlockId.Hash, req.Txs[i], sortKey, prevSortKeys)

		txSize := protoTxSize(txBytes)
		if size+txSize > req.MaxTxBytes {
			break
		}
		result = append(result, txBytes)
		size += txSize
	}

	ctx.Logger().
		With("height", req.Height).
		With("number of txs", len(result)).
		With("size of txs", size).
		With("max size", req.MaxTxBytes).
		With("execution time", time.Since(now).Milliseconds()).
		Info("Prepared transactions")

	return abci.ResponsePrepareProposal{Txs: result}
}

// Creates a transaction with an Arweave block
func (h *prepareProposalHandler) createArweaveTx(ctx sdk.Context, nextArweaveBlock *types.NextArweaveBlock, prevSortKeys *PrevSortKeys) []byte {
	msg := &types.MsgArweaveBlock{
		BlockInfo:    nextArweaveBlock.BlockInfo,
		Transactions: prepareTransactions(nextArweaveBlock.Transactions, prevSortKeys),
	}

	txBuilder := h.txConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msg)
	if err != nil {
		panic(err)
	}

	tx := txBuilder.GetTx()
	bz, err := h.txConfig.TxEncoder()(tx)
	if err != nil {
		panic(err)
	}
	return bz
}

// Sets the PrevSortKey and random values for transactions from the Arweave block
func prepareTransactions(txs []*types.ArweaveTransaction, prevSortKeys *PrevSortKeys) []*types.ArweaveTransactionWithInfo {
	result := make([]*types.ArweaveTransactionWithInfo, len(txs))
	for i, tx := range txs {
		result[i] = &types.ArweaveTransactionWithInfo{
			Transaction: tx,
			PrevSortKey: prevSortKeys.getAndStorePrevSortKey(tx.Contract, tx.SortKey),
			Random:      generateRandomL1(tx.SortKey),
		}
	}
	return result
}

// Sets 'SortKey', 'PrevSortKey' and random value if the transaction is an L2 interaction.
// Returns the original transaction otherwise.
func (h *prepareProposalHandler) prepareDataItem(sequencerBlockHash []byte, txBytes []byte, sortKey *SortKey, prevSortKeys *PrevSortKeys) []byte {
	// decode tx
	tx, err := h.txConfig.TxDecoder()(txBytes)
	if err != nil {
		panic(err)
	}
	dataItem, err := ante.GetL2Interaction(tx)
	if err != nil {
		panic(err)
	}
	if dataItem == nil {
		// Not a data item
		return txBytes
	}

	dataItem.SortKey = sortKey.GetNextValue()

	// Set prev sort key
	contract, err := dataItem.GetContractFromTags()
	if err != nil {
		panic(err)
	}
	dataItem.PrevSortKey = prevSortKeys.getAndStorePrevSortKey(contract, dataItem.SortKey)

	// Set random value
	dataItem.Random = generateRandomL2(sequencerBlockHash, dataItem.SortKey)

	// Encode tx
	txBuilder, err := h.txConfig.WrapTxBuilder(tx)
	if err != nil {
		panic(err)
	}

	err = txBuilder.SetMsgs(dataItem)
	if err != nil {
		panic(err)
	}

	out, err := h.txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		panic(err)
	}

	return out
}

// The transaction size after encoding using Protobuf.
// See: https://pkg.go.dev/github.com/cometbft/cometbft/proto/tendermint/types#Data.Size
func protoTxSize(tx []byte) int64 {
	length := len(tx)
	return 1 + int64(length) + varIntSize(uint64(length))
}

func varIntSize(x uint64) int64 {
	return (int64(math_bits.Len64(x|1)) + 6) / 7
}
