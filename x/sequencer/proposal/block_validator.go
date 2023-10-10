package proposal

import (
	"runtime"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/controller"
	"github.com/warp-contracts/sequencer/x/sequencer/keeper"

	"github.com/warp-contracts/syncer/src/utils/task"
)

type Block struct {
	ctx sdk.Context
	txs []sdk.Tx
}

// Validates transactions in the block proposed by the Proposer.
// Where required, validation is done sequentially.
// In other cases, validation occurs in separate threads per transaction.
type BlockValidator struct {
	*task.Task

	keeper     *keeper.Keeper
	controller controller.ArweaveBlocksController
	input      chan *Block
	output     chan error
}

func NewBlockValidator(keeper *keeper.Keeper, controller controller.ArweaveBlocksController) *BlockValidator {
	validator := new(BlockValidator)
	validator.keeper = keeper
	validator.controller = controller
	validator.input = make(chan *Block)
	validator.output = make(chan error)

	validator.Task = task.NewTask(nil, "block-validator").
		WithSubtaskFunc(validator.run).
		WithWorkerPool(runtime.NumCPU(), 1).
		WithOnAfterStop(func() {
			close(validator.input)
			close(validator.output)
		})

	return validator
}

func (v *BlockValidator) run() error {
	for block := range v.input {
		if len(block.txs) == 0 {
			v.output <- nil
			continue
		}

		txValidator := newTxValidator(block.ctx, v.keeper, v.controller)
		result := newValidationResult(v.output)
		wg := &sync.WaitGroup{}
		wg.Add(1 + len(block.txs)) // one sequential and for each transaction

		v.validateSequentially(txValidator, block.txs, result, wg)
		v.validateInParallel(txValidator, block.txs, result, wg)

		wg.Wait()
		result.sendIfValid()
	}

	return nil
}

func (v *BlockValidator) validateSequentially(txValidator *TxValidator, txs []sdk.Tx, result *validationResult, wg *sync.WaitGroup) {
	v.SubmitToWorker(func() {
		for txIndex, tx := range txs {
			if result.isValid() {
				err := txValidator.validateSequentially(txIndex, tx)
				if err != nil {
					result.sendFirstError(err)
					break
				}
			}
		}
		wg.Done()
	})
}

func (v *BlockValidator) validateInParallel(txValidator *TxValidator, txs []sdk.Tx, result *validationResult, wg *sync.WaitGroup) {
	for txIndex, tx := range txs {
		v.SubmitToWorker(func() {
			if result.isValid() {
				err := txValidator.validateInParallel(txIndex, tx)
				result.sendFirstError(err)
			}
			wg.Done()
		})
	}
}

func (v *BlockValidator) validateBlock(block *Block) error {
	if v == nil || v.IsStopping.Load() {
		return nil
	}

	v.input <- block

	select {
	case <-v.Ctx.Done():
		return nil
	case <-block.ctx.Done():
		return nil
	case err := <-v.output:
		return err
	}
}

func (v *BlockValidator) StopWait() {
	if v != nil && !v.IsStopping.Load() {
		v.Task.StopWait()
	}
}
