package prepare

import (
	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/ante"
	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Sets the 'SortKey' for all L2 interactions and adds an Arweave block transaction to the beginning of the block if needed.
func NewPrepareProposalHandler(keeper keeper.Keeper, arweaveController controller.ArweaveBlocksController, txConfig client.TxConfig) sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req abci.RequestPrepareProposal) abci.ResponsePrepareProposal {
		lastBlock := keeper.MustGetLastArweaveBlock(ctx)
		arweaveHeight := lastBlock.ArweaveBlock.Height
		sequencerHeight := ctx.BlockHeader().Height
		nextBlock := arweaveController.GetNextArweaveBlock(arweaveHeight + 1)
		arweaveBlockTx, i := createArweaveTx(ctx, txConfig, nextBlock)
		sortKey := types.NewSortKey(arweaveHeight + uint64(i), sequencerHeight)

		result := make([][]byte, len(req.Txs) + i)
		if arweaveBlockTx != nil {
			result[0] = arweaveBlockTx
		}

		var size int64 = 0
		txCount := i
		for txCount < len(req.Txs) + i {
			txBytes := setSortKeyInDataItem(txConfig, req.Txs[txCount-i], sortKey)
			size += int64(len(txBytes))
			if size > req.MaxTxBytes {
				break
			}
			result[txCount] = txBytes
			txCount++
		}
		return abci.ResponsePrepareProposal{Txs: result[:txCount]}
	}
}

// Returns the transaction with an Arweave block if it is older than an hour and has not been added to the blockchain yet. 
// Additionally, it returns 1 if such a block exists and 0 otherwise.
func createArweaveTx(ctx sdk.Context, txConfig client.TxConfig, nextArweaveBlock *types.NextArweaveBlock) ([]byte, int) {
	if nextArweaveBlock == nil || !types.IsArweaveBlockOldEnough(ctx, nextArweaveBlock.BlockInfo) {
		return nil, 0
	}

	msg := &types.MsgArweaveBlock{
		BlockInfo:    nextArweaveBlock.BlockInfo,
		Transactions: nextArweaveBlock.Transactions,
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

// Sets the 'SortKey' if the transaction is an L2 interaction. 
// Returns the original transaction otherwise.
func setSortKeyInDataItem(txConfig client.TxConfig, txBytes []byte, sortKey *types.SortKey) []byte {
	tx, err := txConfig.TxDecoder()(txBytes)
	if err != nil {
		panic(err)
	}
	dataItem, err := ante.GetL2Interaction(tx)
	if err != nil {
		panic(err)
	}
	if dataItem != nil {
		dataItem.SortKey = sortKey.GetNextValue()
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
