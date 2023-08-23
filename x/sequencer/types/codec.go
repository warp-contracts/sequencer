package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	cryptocodec "github.com/warp-contracts/sequencer/crypto/codec"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDataItem{}, "sequencer/DataItem", nil)
	cdc.RegisterConcrete(&MsgArweaveBlockInfo{}, "sequencer/LastArweaveBlock", nil)
	cdc.RegisterConcrete(&MsgArweaveTransaction{}, "sequencer/ArweaveTransaction", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDataItem{},
	)
	cryptocodec.RegisterInterfaces(registry)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgArweaveBlockInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgArweaveTransaction{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
