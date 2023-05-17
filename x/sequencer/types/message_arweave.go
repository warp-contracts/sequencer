package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgArweave = "arweave"

var _ sdk.Msg = &MsgArweave{}

func NewMsgArweave(creator string, format string, id string, lastTx string, owner string, tags string, target string, quantity string, dataRoot string, dataSize string, data string, reward string, signature string) *MsgArweave {
	return &MsgArweave{
		Creator:   creator,
		Format:    format,
		Id:        id,
		LastTx:    lastTx,
		Owner:     owner,
		Tags:      tags,
		Target:    target,
		Quantity:  quantity,
		DataRoot:  dataRoot,
		DataSize:  dataSize,
		Data:      data,
		Reward:    reward,
		Signature: signature,
	}
}

func (msg *MsgArweave) Route() string {
	return RouterKey
}

func (msg *MsgArweave) Type() string {
	return TypeMsgArweave
}

func (msg *MsgArweave) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgArweave) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgArweave) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
