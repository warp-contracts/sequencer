package ante

import (
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

// Stores identifiers of L2 interactions contained in the last sequencer block
type BlockInteractions struct {
	dataItems map[string]bool
}

func NewBlockInteractions() *BlockInteractions {
	return &BlockInteractions{
		dataItems: make(map[string]bool),
	}
}

func (bi *BlockInteractions) NewBlock() {
	clear(bi.dataItems)
}

func (bi *BlockInteractions) Contains(dataItem *types.MsgDataItem) bool {
	dataItemId := dataItem.DataItem.Id.Base64()
	return bi.dataItems[dataItemId]
}

func (bi *BlockInteractions) Add(dataItem *types.MsgDataItem) {
	dataItemId := dataItem.DataItem.Id.Base64()
	bi.dataItems[dataItemId] = true
}
