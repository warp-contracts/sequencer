package prepare

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// If there is an Arweave block that is over an hour old and has not yet been added to the blockchain,
// a transaction containing such a block is added to the beginning of the Cosmos block proposal.
func NewPrepareProposalHandler(keeper keeper.Keeper, arweaveController controller.ArweaveBlocksController, txConfig client.TxConfig, logger log.Logger) sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req abci.RequestPrepareProposal) abci.ResponsePrepareProposal {
		txs := req.Txs
		lastBlock := keeper.MustGetLastArweaveBlock(ctx)
		nextBlock := arweaveController.GetNextArweaveBlock(lastBlock.Height + 1)

		if nextBlock != nil && types.IsArweaveBlockOldEnough(ctx, nextBlock.BlockInfo) {
			txs = append([][]byte{createArweaveTx(ctx, txConfig, nextBlock)}, txs...)
		}

		return abci.ResponsePrepareProposal{Txs: txs}
	}
}

func createArweaveTx(ctx sdk.Context, txConfig client.TxConfig, nextArweaveBlock *types.NextArweaveBlock) []byte {
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
	return bz
}
