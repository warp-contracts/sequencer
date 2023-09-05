package controller

import (
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type ArweaveBlocksControllerMock struct {
	block *types.NextArweaveBlock
}

func (mock ArweaveBlocksControllerMock) StartController(height uint64) {
}

func (mock ArweaveBlocksControllerMock) IsRunning() bool {
	return true
}

func (mock ArweaveBlocksControllerMock) GetNextArweaveBlock(height uint64) *types.NextArweaveBlock {
	return mock.block
}

func (mock ArweaveBlocksControllerMock) RemoveNextArweaveBlocksUpToHeight(height uint64) {
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
