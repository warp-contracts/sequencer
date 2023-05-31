package ante

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/utils/arweave"
	"github.com/warp-contracts/syncer/src/utils/bundlr"
	"github.com/warp-contracts/syncer/src/utils/tool"
)

func newTxBuilder() client.TxBuilder {
	return simapp.MakeTestEncodingConfig().TxConfig.NewTxBuilder()
}

func exampleDataItem() types.MsgDataItem {
	dataItem := bundlr.BundleItem{
		SignatureType: 1,
		Signature: arweave.Base64String(tool.RandomString(32)),
		Target:        arweave.Base64String(tool.RandomString(32)),
		Anchor: arweave.Base64String(tool.RandomString(32)),
		Tags: bundlr.Tags{bundlr.Tag{Name: "1", Value: "2"}, bundlr.Tag{Name: "3", Value: "4"}},
		Data: arweave.Base64String(tool.RandomString(100)),
	}

	return types.MsgDataItem{
		Creator:  "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
		DataItem: dataItem,
	}
}
