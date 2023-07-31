package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateLastArweaveBlock = "create_last_arweave_block"
	TypeMsgUpdateLastArweaveBlock = "update_last_arweave_block"
	TypeMsgDeleteLastArweaveBlock = "delete_last_arweave_block"
)

var _ sdk.Msg = &MsgCreateLastArweaveBlock{}

func NewMsgCreateLastArweaveBlock(creator string) *MsgCreateLastArweaveBlock {
	return &MsgCreateLastArweaveBlock{
		Creator: creator,
	}
}

func (msg *MsgCreateLastArweaveBlock) Route() string {
	return RouterKey
}

func (msg *MsgCreateLastArweaveBlock) Type() string {
	return TypeMsgCreateLastArweaveBlock
}

func (msg *MsgCreateLastArweaveBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLastArweaveBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLastArweaveBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateLastArweaveBlock{}

func NewMsgUpdateLastArweaveBlock(creator string) *MsgUpdateLastArweaveBlock {
	return &MsgUpdateLastArweaveBlock{
		Creator: creator,
	}
}

func (msg *MsgUpdateLastArweaveBlock) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLastArweaveBlock) Type() string {
	return TypeMsgUpdateLastArweaveBlock
}

func (msg *MsgUpdateLastArweaveBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLastArweaveBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLastArweaveBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLastArweaveBlock{}

func NewMsgDeleteLastArweaveBlock(creator string) *MsgDeleteLastArweaveBlock {
	return &MsgDeleteLastArweaveBlock{
		Creator: creator,
	}
}
func (msg *MsgDeleteLastArweaveBlock) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLastArweaveBlock) Type() string {
	return TypeMsgDeleteLastArweaveBlock
}

func (msg *MsgDeleteLastArweaveBlock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLastArweaveBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLastArweaveBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
