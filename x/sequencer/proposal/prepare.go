package proposal

import (
	"fmt"
	math_bits "math/bits"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/ante"
	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Sets sort keys for all L2 interactions and adds an Arweave block transaction to the beginning of the block if needed.
func NewPrepareProposalHandler(keeper *keeper.Keeper, arweaveController controller.ArweaveBlocksController, txConfig client.TxConfig) sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req abci.RequestPrepareProposal) abci.ResponsePrepareProposal {
		lastBlock := keeper.MustGetLastArweaveBlock(ctx)
		arweaveHeight := lastBlock.ArweaveBlock.Height
		sequencerHeight := ctx.BlockHeight()
		nextBlock := arweaveController.GetNextArweaveBlock(arweaveHeight + 1)
		lastSortKeys := newLastSortKeys(keeper, ctx)
		arweaveBlockTx, i := createArweaveTx(ctx, txConfig, nextBlock, lastSortKeys)
		sortKey := newSortKey(arweaveHeight+uint64(i), sequencerHeight)
		var size int64 = 0

		result := make([][]byte, len(req.Txs)+i)
		if arweaveBlockTx != nil {
			result[0] = arweaveBlockTx
			size += protoTxSize(arweaveBlockTx)
			if size > req.MaxTxBytes {
				panic(fmt.Sprintf("MaxTxBytes limit (%d) is too small! It is smaller than the size of the Arweave block (%d)",
					req.MaxTxBytes, size))
			}
		}

		txCount := i
		for txCount < len(req.Txs)+i {
			txBytes := setSortKeysInDataItem(txConfig, req.Txs[txCount-i], sortKey, lastSortKeys)
			txSize := protoTxSize(txBytes)
			if size+txSize > req.MaxTxBytes {
				break
			}
			result[txCount] = txBytes
			txCount++
			size += txSize
		}

		ctx.Logger().
			With("height", req.Height).
			With("number of txs", txCount).
			With("size of txs", size).
			With("max size", req.MaxTxBytes).
			Info("Prepared transactions")

		return abci.ResponsePrepareProposal{Txs: result[:txCount]}
	}
}

// Returns the transaction with an Arweave block if it is older than an hour and has not been added to the blockchain yet.
// Additionally, it returns 1 if such a block exists and 0 otherwise.
func createArweaveTx(ctx sdk.Context, txConfig client.TxConfig, nextArweaveBlock *types.NextArweaveBlock, lastSortKeys *LastSortKeys) ([]byte, int) {
	if nextArweaveBlock == nil || !types.IsArweaveBlockOldEnough(ctx, nextArweaveBlock.BlockInfo) {
		return nil, 0
	}

	msg := &types.MsgArweaveBlock{
		BlockInfo:    nextArweaveBlock.BlockInfo,
		Transactions: prepareTransactions(nextArweaveBlock.Transactions, lastSortKeys),
	}

	txBuilder := txConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msg)
	if err != nil {
		panic(err)
	}

	tx := txBuilder.GetTx()
	bz, err := txConfig.TxEncoder()(tx)
	if err != nil {
		panic(err)
	}
	return bz, 1
}

// Sets the LastSortKey values for transactions from the Arweave block
func prepareTransactions(txs []*types.ArweaveTransaction, lastSortKeys *LastSortKeys) []*types.ArweaveTransactionWithLastSortKey {
	result := make([]*types.ArweaveTransactionWithLastSortKey, len(txs))
	for i, tx := range txs {
		result[i] = &types.ArweaveTransactionWithLastSortKey{
			Transaction: tx,
			LastSortKey: lastSortKeys.getAndStoreLastSortKey(tx.Contract, tx.SortKey),
		}
	}
	return result
}

// Sets 'SortKey' and 'LastSortKey' if the transaction is an L2 interaction.
// Returns the original transaction otherwise.
func setSortKeysInDataItem(txConfig client.TxConfig, txBytes []byte, sortKey *SortKey, lastSortKeys *LastSortKeys) []byte {
	// decode tx
	tx, err := txConfig.TxDecoder()(txBytes)
	if err != nil {
		panic(err)
	}
	dataItem, err := ante.GetL2Interaction(tx)
	if err != nil {
		panic(err)
	}
	if dataItem != nil {
		// set sort key
		dataItem.SortKey = sortKey.GetNextValue()

		// set last sort key
		contract, err := dataItem.GetContractFromTags()
		if err != nil {
			panic(err)
		}
		dataItem.LastSortKey = lastSortKeys.getAndStoreLastSortKey(contract, dataItem.SortKey)

		// encode tx
		txBuilder, err := txConfig.WrapTxBuilder(tx)
		if err != nil {
			panic(err)
		}
		err = txBuilder.SetMsgs(dataItem)
		if err != nil {
			panic(err)
		}
		txBytes, err = txConfig.TxEncoder()(txBuilder.GetTx())
		if err != nil {
			panic(err)
		}
	}
	return txBytes
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
