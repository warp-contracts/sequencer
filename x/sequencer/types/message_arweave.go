package types

import (
	"encoding/json"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/warp-contracts/syncer/src/utils/arweave"
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
	// Ensure message contains an Arweave transaction
	tx, err := msg.ToArweaveTx()
	if err != nil {
		return err
	}

	err = tx.Verify()
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgArweave) ToArweaveTx() (out *arweave.Transaction, err error) {
	out = &arweave.Transaction{
		Format:   int(msg.Format),
		ID:       msg.Id,
		Quantity: msg.Quantity,
		DataSize: msg.DataSize,
		Reward:   msg.Reward,
	}

	err = out.LastTx.Decode(msg.LastTx)
	if err != nil {
		return
	}
	err = out.Owner.Decode(msg.Owner)
	if err != nil {
		return
	}
	err = out.Target.Decode(msg.Target)
	if err != nil {
		return
	}
	err = out.Data.Decode(msg.Data)
	if err != nil {
		return
	}
	err = out.DataRoot.Decode(msg.DataRoot)
	if err != nil {
		return
	}
	err = out.Signature.Decode(msg.Signature)
	if err != nil {
		return
	}

	out.Tags = make([]arweave.Tag, 0, len(msg.Tags))
	for _, tag := range msg.Tags {
		var tmp arweave.Tag
		err = tmp.Name.Decode(tag.Name)
		if err != nil {
			return
		}
		err = tmp.Value.Decode(tag.Value)
		if err != nil {
			return
		}
		out.Tags = append(out.Tags, tmp)
	}

	return
}
