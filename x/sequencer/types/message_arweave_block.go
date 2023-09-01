package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgArweaveBlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgArweaveBlock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.BlockInfo.Hash) != 32 {
		return errors.Wrapf(ErrBadArweaveHashLength, "hash length should be 32 and is %d", len(msg.BlockInfo.Hash))
	}

	if msg.BlockInfo.Height < 1231216 {
		return errors.Wrap(ErrBadArweaveHeight, "block height should be greater than 1231215")
	}

	if msg.BlockInfo.Timestamp < 1690809540 {
		return errors.Wrap(ErrBadArweaveTimestamp, "block timestamp should be greater than 1690809539")
	}

	for _, transaction := range msg.Transactions {
		if err := transaction.Transaction.Verify(); err != nil {
			return err
		}
	}

	return nil
}
