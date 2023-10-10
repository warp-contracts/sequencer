package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgArweaveBlock = "arweave_block"
)

var _ sdk.Msg = &MsgArweaveBlock{}

func (msg *MsgArweaveBlock) Route() string {
	return RouterKey
}

func (msg *MsgArweaveBlock) Type() string {
	return TypeMsgArweaveBlock
}

func (msg *MsgArweaveBlock) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgArweaveBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgArweaveBlock) ValidateBasic() error {
	if len(msg.BlockInfo.Hash) <= 0 {
		return errors.Wrapf(ErrBadArweaveHash, "hash length should be greater than zero")
	}

	if msg.BlockInfo.Height <= 0 {
		return errors.Wrap(ErrBadArweaveHeight, "block height should be greater than zero")
	}

	if msg.BlockInfo.Timestamp <= 0 {
		return errors.Wrap(ErrBadArweaveTimestamp, "block timestamp should be greater than zero")
	}

	return nil
}
