package ante

import (
	"fmt"

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Stores identifiers of L2 interactions contained in the last sequencer block
type BlockInteractions struct {
	height    int64
	dataItems map[string]bool
}

func NewBlockInteractions() *BlockInteractions {
	return &BlockInteractions{
		dataItems: make(map[string]bool),
	}
}

func (bi *BlockInteractions) NewBlock(height int64) {
	bi.checkHeight(true, height - 1)

	bi.height = height
	clear(bi.dataItems)
}

func (bi *BlockInteractions) Contains(height int64, dataItem *types.MsgDataItem) bool {
	bi.checkHeight(false, height)

	dataItemId := dataItem.DataItem.Id.Base64()
	return bi.dataItems[dataItemId]
}

func (bi *BlockInteractions) Add(height int64, dataItem *types.MsgDataItem) {
	bi.checkHeight(false, height)

	dataItemId := dataItem.DataItem.Id.Base64()
	bi.dataItems[dataItemId] = true
}

func (bi *BlockInteractions) checkHeight(init bool, height int64) {
	if bi.height == 0 {
		if !init {
			panic("BlockInteractions has not been initialized")
		}
	} else if bi.height != height {
		panic(
			fmt.Sprintf("Inconsistency in block height: saved interactions for height %d, processed height %d",
				bi.height, height),
		)
	}
}
