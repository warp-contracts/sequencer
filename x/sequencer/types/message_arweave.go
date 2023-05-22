package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

const TypeMsgArweave = "arweave"

var _ sdk.Msg = &MsgArweave{}

func NewMsgArweave(creator string, format string, id string, lastTx string, owner string, tagsJson string, target string, quantity string, dataRoot string, dataSize string, data string, reward string, signature string) (*MsgArweave, error) {
	formatNumber, err := strconv.ParseInt(format, 10, 32)
	if err != nil {
		return nil, err
	}
	tags, err := parseTags(tagsJson)
	if err != nil {
		return nil, err
	}
	return &MsgArweave{
		Creator:   creator,
		Format:    int32(formatNumber),
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
	}, nil
}

func parseTags(tagsJson string) ([]*Tag, error) {
	var tagsMap map[string]string
	err := json.Unmarshal([]byte(tagsJson), &tagsMap)
	if err != nil {
		return nil, err
	}
	var tags []*Tag
	for name, value := range tagsMap {
		tag := Tag{Name: name, Value: value}
		tags = append(tags, &tag)
	}
	return tags, nil
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
