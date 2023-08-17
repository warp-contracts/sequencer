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

func NewMsgLastArweaveBlock(creator string, height uint64, timestamp uint64, hash []byte) *MsgLastArweaveBlock {
	return &MsgLastArweaveBlock{
		Creator: creator,
		Height: height,
		Timestamp: timestamp,
		Hash: hash,
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
