package types

import (
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDataItem = "data_item"

var _ sdk.Msg = &MsgDataItem{}

func NewMsgDataItem(creator string, data string) (out *MsgDataItem, err error) {
	out = &MsgDataItem{
		Creator: creator,
	}

	dataItemBytes, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return
	}

	// Data item is base64url encoded
	err = out.DataItem.Unmarshal(dataItemBytes)
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
	_, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Verifies DataItem acording to the ANS-104 standard. Verifies signature.
	// https://github.com/ArweaveTeam/arweave-standards/blob/master/ans/ANS-104.md#21-verifying-a-dataitem
	err = msg.DataItem.Verify()
	if err != nil {
		return
	}

	return nil
}
