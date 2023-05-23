package types

import (
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDataItem = "arweave"

var _ sdk.Msg = &MsgDataItem{}

func NewMsgDataItem(creator string, data string) (out *MsgDataItem, err error) {
	out = &MsgDataItem{
		Creator: creator,
	}

	// Data item is base64url encoded
	out.DataItem, err = base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return
	}

	return
}

func (msg *MsgDataItem) Route() string {
	return RouterKey
}

func (msg *MsgDataItem) Type() string {
	return TypeMsgDataItem
}

func (msg *MsgDataItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDataItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDataItem) ValidateBasic() (err error) {
	// Ensure data item is in the correct format

	_, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
