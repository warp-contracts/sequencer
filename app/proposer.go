package app

import (
	"errors"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	sequencermodulekeeper "github.com/warp-contracts/sequencer/x/sequencer/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// https://docs.cosmos.network/main/building-apps/app-mempool
type Proposer struct {
	txConfig   client.TxConfig
	decoder    sdk.TxDecoder
	mempool    mempool.Mempool
	txVerifier baseapp.ProposalTxVerifier
	keeper     sequencermodulekeeper.Keeper
}

func NewProposer(txConfig client.TxConfig, mempool mempool.Mempool, txVerifier baseapp.ProposalTxVerifier, keeper sequencermodulekeeper.Keeper) *Proposer {
	return &Proposer{
		txConfig:   txConfig,
		decoder:    txConfig.TxDecoder(),
		mempool:    mempool,
		txVerifier: txVerifier,
		keeper:     keeper,
	}
}

func (self Proposer) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req abci.RequestPrepareProposal) abci.ResponsePrepareProposal {
		// If the mempool is nil or NoOp we simply return the transactions
		// requested from CometBFT, which, by default, should be in FIFO order.
		_, isNoOp := self.mempool.(mempool.NoOpMempool)
		if self.mempool == nil || isNoOp {
			return abci.ResponsePrepareProposal{Txs: req.Txs}
		}

		var (
			selectedTxs  [][]byte
			totalTxBytes int64
			err          error
		)

		iterator := self.mempool.Select(ctx, req.Txs)

		idx := 0
		for iterator != nil {
			memTx := iterator.Tx()

			for _, msg := range memTx.GetMsgs() {
				dataItem, isDataItem := msg.(*types.MsgDataItem)
				if !isDataItem {
					continue
				}

				block, found := self.keeper.GetLastArweaveBlock(ctx)
				if !found {
					// FIXME: Handle no arweave block
					block = types.LastArweaveBlock{
						Height: 12,
					}
					// panic("last arweave block not found")
				}

				// Old sort key format <12char of arweave block height>,<13char timestamp>,<64char arweave block hash>
				// v1 sort keys must always come before v2 sort keys
				// 000001236585,0000000000000,5517c5c6ed15f873f9be754d899fc15314f01e8fbd8a04fee8d2268c00f506d5

				// FIXME: Go through padding
				dataItem.SortKey = fmt.Sprintf("%012d,%020d,%010d", block.Height, req.Height, idx)

				memTx, err = types.CreateTxWithDataItem(self.txConfig, *dataItem)
				if err != nil {
					panic(err)
				}

				break

			}

			bz, err := self.txVerifier.PrepareProposalVerifyTx(memTx)
			if err != nil {
				err := self.mempool.Remove(memTx)
				if err != nil && !errors.Is(err, mempool.ErrTxNotFound) {
					panic(err)
				}
			} else {
				txSize := int64(len(bz))
				if totalTxBytes += txSize; totalTxBytes <= req.MaxTxBytes {
					selectedTxs = append(selectedTxs, bz)
				} else {
					// We've reached capacity per req.MaxTxBytes so we cannot select any
					// more transactions.
					break
				}
			}

			iterator = iterator.Next()
			idx++
		}

		return abci.ResponsePrepareProposal{Txs: selectedTxs}

		// Fill in sort key for each DataItem message
		// txs := ctypes.ToTxs(req.Txs)
		// for _, txBytes := range txs {
		// 	tx, err := self.decoder(txBytes)
		// 	if err != nil {
		// 		goto end
		// 	}

		// 	for _, msg := range tx.GetMsgs() {
		// 		if proto.MessageName(msg) == "sequencer.sequencer.MsgDataItem" {
		// 			ctx.Logger().Info("--------------------------------------------------------")

		// 		}
		// 	}

		// }

		// end:
		// 	// FIXME: Check if txs aren't too big after adding the Arweave block https://github.com/cometbft/cometbft/issues/980
		// 	return abci.ResponsePrepareProposal{
		// 		Txs: req.Txs,
		// 	}
	}
}
