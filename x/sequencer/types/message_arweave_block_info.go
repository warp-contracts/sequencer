package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgArweaveBlockInfo = "arweave_block_info"
)

var _ sdk.Msg = &MsgArweaveBlockInfo{}

func NewMsgArweaveBlockInfo(creator string, height uint64, timestamp uint64, hash []byte) *MsgArweaveBlockInfo {
	return &MsgArweaveBlockInfo{
		Creator:   creator,
		Height:    height,
		Timestamp: timestamp,
		Hash:      hash,
	}
}

func (msg *MsgArweaveBlockInfo) Route() string {
	return RouterKey
}

func (msg *MsgArweaveBlockInfo) Type() string {
	return TypeMsgArweaveBlockInfo
}

func (msg *MsgArweaveBlockInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgArweaveBlockInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgArweaveBlockInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Hash) != 32 {
		return errors.Wrapf(ErrBadArweaveHashLength, "hash length should be 32 and is %d", len(msg.Hash))
	}

	if msg.Height < 1231216 {
		return errors.Wrap(ErrBadArweaveHeight, "block height should be greater than 1231215")
	}

	if msg.Timestamp < 1690809540 {
		return errors.Wrap(ErrBadArweaveTimestamp, "block timestamp should be greater than 1690809539")
	}

	return nil
}
