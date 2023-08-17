package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgLastArweaveBlock = "last_arweave_block"
)

var _ sdk.Msg = &MsgLastArweaveBlock{}

func NewMsgLastArweaveBlock(creator string) *MsgLastArweaveBlock {
	return &MsgLastArweaveBlock{
		Creator: creator,
	}
}

func (msg *MsgLastArweaveBlock) Route() string {
	return RouterKey
}

func (msg *MsgLastArweaveBlock) Type() string {
	return TypeMsgLastArweaveBlock
}

func (msg *MsgLastArweaveBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLastArweaveBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLastArweaveBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
