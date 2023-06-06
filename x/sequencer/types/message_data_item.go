package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	keys "github.com/warp-contracts/sequencer/crypto/keys/arweave"
)

const TypeMsgDataItem = "data_item"
// TODO: move to syncer/src/utils/warp/tags.go
const SequencerNonceTag = "Sequencer-Nonce"

var _ sdk.Msg = &MsgDataItem{}

func (msg *MsgDataItem) Route() string {
	return RouterKey
}

func (msg *MsgDataItem) Type() string {
	return TypeMsgDataItem
}

func (msg *MsgDataItem) GetCreator() sdk.AccAddress {
	pubKey := keys.UnmarshalPubkey(msg.DataItem.Owner)
	return pubKey.AccAddress()
}

func (msg *MsgDataItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.GetCreator()}
}

func (msg *MsgDataItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDataItem) ValidateBasic() (err error) {
	// Verifies DataItem acording to the ANS-104 standard. Verifies signature.
	// https://github.com/ArweaveTeam/arweave-standards/blob/master/ans/ANS-104.md#21-verifying-a-dataitem
	err = msg.DataItem.Verify()
	if err != nil {
		return
	}

	err = msg.DataItem.VerifySignature()
	if err != nil {
		return
	}

	return nil
}

func (msg *MsgDataItem) GetSequenceFromTags() (uint64, error) {
	for _, tag := range msg.DataItem.Tags {
		if tag.Name == SequencerNonceTag {
			return strconv.ParseUint(tag.Value, 10, 64)
		}
	}
	return 0, sdkerrors.Wrapf(ErrNoSequencerNonceTag, "data item does not have \"%s\" tag", SequencerNonceTag)
}
