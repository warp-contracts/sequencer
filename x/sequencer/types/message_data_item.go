package types

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"

	"cosmossdk.io/errors"
	"cosmossdk.io/x/tx/signing"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/api/sequencer/sequencer"
	"github.com/warp-contracts/sequencer/crypto/keys/arweave"
	"github.com/warp-contracts/sequencer/crypto/keys/ethereum"

	"github.com/warp-contracts/syncer/src/utils/bundlr"
	"github.com/warp-contracts/syncer/src/utils/smartweave"
	"github.com/warp-contracts/syncer/src/utils/warp"
)

var _ sdk.Msg = &MsgDataItem{}

func (msg *MsgDataItem) GetSenderAddress() sdk.AccAddress {
	pubKey, err := msg.GetPublicKey()
	if err != nil {
		panic(err)
	}
	return sdk.AccAddress(pubKey.Address())
}

func (msg *MsgDataItem) ValidateBasic() (err error) {
	// Data item validation is done in the `AnteHandler` and `processProposalHandler` using the `Verify` method
	// to have greater control over it and avoid executing it during recheckTx.
	return nil
}

func (msg *MsgDataItem) Verify() (err error) {
	// Verifies DataItem acording to the ANS-104 standard. Verifies signature.
	// https://github.com/ArweaveTeam/arweave-standards/blob/master/ans/ANS-104.md#21-verifying-a-dataitem
	err = msg.DataItem.Verify()
	if err != nil {
		return
	}
	return msg.DataItem.VerifySignature()
}

func (msg *MsgDataItem) GetNonceFromTags() (uint64, error) {
	nonce, found := msg.DataItem.GetTag(warp.TagSequencerNonce)
	if found {
		return strconv.ParseUint(nonce, 10, 64)
	}
	return 0, errors.Wrapf(ErrNoSequencerNonceTag, "data item does not have \"%s\" tag", warp.TagSequencerNonce)
}

func (msg *MsgDataItem) GetContractFromTags() (string, error) {
	contract, found := msg.DataItem.GetTag(smartweave.TagContractTxId)
	if found {
		return contract, nil
	}
	return "", errors.Wrapf(ErrNoContractTag, "data item does not have \"%s\" tag", smartweave.TagContractTxId)
}

func (msg *MsgDataItem) GetPublicKey() (cryptotypes.PubKey, error) {
	return GetPublicKey(msg.DataItem.SignatureType, msg.DataItem.Owner)
}

func GetPublicKey(signatureType bundlr.SignatureType, owner []byte) (cryptotypes.PubKey, error) {
	switch signatureType {
	case bundlr.SignatureTypeArweave:
		pubKey := arweave.FromOwner(owner)
		return pubKey, nil
	case bundlr.SignatureTypeEthereum:
		pubKey, err := ethereum.FromOwner(owner)
		return pubKey, err
	default:
		return nil, bundlr.ErrUnsupportedSignatureType
	}
}

type DataItemInfo struct {
	DataItemId string `json:"data_item_id"`
	Nonce      uint64 `json:"nonce"`
	Sender     string `json:"sender"`
}

func (msg *MsgDataItem) GetInfo() DataItemInfo {
	nonce, _ := msg.GetNonceFromTags()
	return DataItemInfo{
		DataItemId: msg.DataItem.Id.Base64(),
		Nonce:      nonce,
		Sender:     msg.GetSenderAddress().String(),
	}
}

// FIXME 
// In the function returning signers, an Unmarshal occurs. 
// Check whether this could be avoided and why the MsgDataItem type from pulsar.go files is used here, 
// and the type with the same name from pb.go in the ante handler.
func ProvideMsgDataItemGetSingers() signing.CustomGetSigner {
	return signing.CustomGetSigner{
		MsgType: proto.MessageName(&sequencer.MsgDataItem{}),
		Fn: func(msg proto.Message) ([][]byte, error) {
			msgDataItem, ok := msg.(*sequencer.MsgDataItem)
			if !ok {
				return nil, fmt.Errorf("Invalid message type: %T", msg)
			}

			bundleItem := new(bundlr.BundleItem)
			err := bundleItem.Unmarshal(msgDataItem.DataItem)
			if err != nil {
				return nil, err
			}
			pubKey, err := GetPublicKey(bundleItem.SignatureType, bundleItem.Owner)
			if err != nil {
				return nil, err
			}
			return [][]byte{pubKey.Address()}, nil
		},
	}
}
