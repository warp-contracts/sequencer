package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	cryptocodec "github.com/warp-contracts/sequencer/crypto/codec"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	cryptocodec.RegisterInterfaces(registry)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDataItem{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgArweaveBlock{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil))
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
