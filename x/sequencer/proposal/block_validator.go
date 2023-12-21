package proposal

import (
	"sync"

	"github.com/cometbft/cometbft/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

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

	logger                        log.Logger
	provider                      *ArweaveBlockProvider
	input                         chan *Block
	output                        chan *InvalidTxError
	consecutiveArweaveBlockErrors int
}

func NewBlockValidator(provider *ArweaveBlockProvider, logger log.Logger) *BlockValidator {
	validator := new(BlockValidator)
	validator.provider = provider
	validator.input = make(chan *Block)
	validator.output = make(chan *InvalidTxError)
	validator.logger = logger.With("module", "block-validator")

	validator.Task = task.NewTask(nil, "block-validator").
		WithSubtaskFunc(validator.run).
		WithWorkerPool(8, 1000).
		// WithWorkerPool(runtime.NumCPU(), 1).
		WithOnAfterStop(func() {
			// We are closing only the `output` channel to avoid panicking when sending to the `input` channel in the `ValidateBlock` method.
			close(validator.output)
		})

	return validator
}

func (v *BlockValidator) run() error {
	for {
		select {
		case <-v.Ctx.Done():
			return nil
		case block := <-v.input:
			txValidator := newTxValidator(block.ctx, v.provider)
			result := newValidationResult(v.output)

			if len(block.txs) == 0 {
				v.validateEmptyBlock(txValidator, result)
			} else {
				wg := &sync.WaitGroup{}
				wg.Add(1 + len(block.txs)) // one sequential and for each transaction

				v.validateSequentially(txValidator, block.txs, result, wg)
				v.validateInParallel(txValidator, block.txs, result, wg)

				wg.Wait()
			}

			result.sendIfNoError()
		}
	}
}

func (v *BlockValidator) validateEmptyBlock(txValidator *TxValidator, result *validationResult) {
	err := txValidator.verifyNoTransactions()
	if err != nil {
		result.sendFirstError(err)
	}
}

func (v *BlockValidator) validateSequentially(txValidator *TxValidator, txs []sdk.Tx, result *validationResult, wg *sync.WaitGroup) {
	v.SubmitToWorker(func() {
		for txIndex, tx := range txs {
			if result.isNotSent() {
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
			if result.isNotSent() {
				err := txValidator.validateInParallel(txIndex, tx)
				result.sendFirstError(err)
			}
			wg.Done()
		})
	}
}

func (v *BlockValidator) ValidateBlock(block *Block) error {
	if v == nil {
		return nil
	}

	v.logger.
		With("height", block.ctx.BlockHeight()).
		With("timestamp", block.ctx.BlockTime()).
		Debug("Validate block")

	// sending the block to the input channel (with checking whether the task is not stopped)
	select {
	case <-v.Ctx.Done():
		return nil
	case <-block.ctx.Done():
		return nil
	case v.input <- block:
	}

	// receiving the validation result from the output channel (with checking whether the task is not stopped)
	select {
	case <-v.Ctx.Done():
		return nil
	case <-block.ctx.Done():
		return nil
	case err := <-v.output:
		return v.handleInvalidTxError(err)
	}
}

func (v *BlockValidator) handleInvalidTxError(err *InvalidTxError) error {
	if err == nil {
		v.consecutiveArweaveBlockErrors = 0
		return nil
	}

	if err.errorType == INVALID_ARWEAVE {
		v.consecutiveArweaveBlockErrors++
		v.logger.
			With("consecutive_errors", v.consecutiveArweaveBlockErrors).
			Debug("Invalid Arweave block error")
		if v.consecutiveArweaveBlockErrors > 10 {
			v.consecutiveArweaveBlockErrors = 0
			v.logger.Error("Controller restart due to too many consecutive Arweave block errors")
			v.provider.controller.Restart()
		}
	}

	return err.err
}

func (v *BlockValidator) StopWait() {
	if v != nil {
		v.Task.StopWait()
	}
}
