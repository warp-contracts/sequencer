package proposal

import (
	"runtime"
	"sync"

	"github.com/cometbft/cometbft/libs/log"

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
	logger     log.Logger
	Input      chan *Block
	Output     chan bool
}

func newBlockValidator(keeper *keeper.Keeper, controller controller.ArweaveBlocksController, logger log.Logger) *BlockValidator {
	validator := new(BlockValidator)
	validator.keeper = keeper
	validator.controller = controller
	validator.logger = logger
	validator.Input = make(chan *Block)
	validator.Output = make(chan bool)

	validator.Task = task.NewTask(controller.GetConfig(), "block-validator").
		WithSubtaskFunc(validator.run).
		WithWorkerPool(runtime.NumCPU(), 1)

	return validator
}

func (v *BlockValidator) run() error {
	for block := range v.Input {
		if len(block.txs) == 0 {
			v.Output <- true
			continue
		}

		txValidator := newTxValidator(block.ctx, v.keeper, v.controller, v.logger)
		result := newValidationResult(v.Output)
		wg := &sync.WaitGroup{}
		wg.Add(len(block.txs))

		v.validateInParallel(txValidator, block.txs, result, wg)
		v.validateSequentially(txValidator, block.txs, result)

		wg.Wait()
		result.sendIfValid()
	}

	return nil
}

func (v *BlockValidator) validateInParallel(txValidator *TxValidator, txs []sdk.Tx, result *validationResult, wg *sync.WaitGroup) {
	for txIndex, tx := range txs {
		v.SubmitToWorker(func() {
			txResult := result.isValid() && txValidator.validateInParallel(txIndex, tx)
			result.sendFirstInvalid(txResult)
			wg.Done()
		})
	}
}

func (v *BlockValidator) validateSequentially(txValidator *TxValidator, txs []sdk.Tx, result *validationResult) {
	for txIndex, tx := range txs {
		txResult := result.isValid() && txValidator.validateSequentially(txIndex, tx)
		result.sendFirstInvalid(txResult)
		if !txResult {
			return
		}
	}
}
