package types

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"cosmossdk.io/errors"
	"cosmossdk.io/x/tx/signing"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgArweaveBlock{}

func NewMsgArweaveBlock(creator string) *MsgArweaveBlock {
	return &MsgArweaveBlock{}
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

func ProvideMsgArweaveBlockGetSingers() signing.CustomGetSigner {
	return signing.CustomGetSigner{
		MsgType: protoreflect.FullName("sequencer.sequencer.MsgArweaveBlock"),
		Fn: nil,
	}
}
