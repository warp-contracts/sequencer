package controller

import (
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type ArweaveBlocksControllerMock struct {
	block *types.NextArweaveBlock
}

func (mock ArweaveBlocksControllerMock) SetLastAcceptedBlock(*types.ArweaveBlockInfo) {
}

func (mock ArweaveBlocksControllerMock) GetNextArweaveBlock(height uint64) *types.NextArweaveBlock {
	return mock.block
}

func (mock ArweaveBlocksControllerMock) StopWait() {
}

func MockArweaveBlocksController(blockInfo *types.ArweaveBlockInfo) ArweaveBlocksController {
	if blockInfo == nil {
		return ArweaveBlocksControllerMock{}
	}
	return ArweaveBlocksControllerMock{
		block: &types.NextArweaveBlock{
			BlockInfo: blockInfo,
		},
	}
}
