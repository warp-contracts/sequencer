package types

import (
	"cosmossdk.io/errors"
	"strconv"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/crypto/keys/arweave"
	"github.com/warp-contracts/sequencer/crypto/keys/ethereum"

	"github.com/warp-contracts/syncer/src/utils/bundlr"
	"github.com/warp-contracts/syncer/src/utils/warp"
	"github.com/warp-contracts/syncer/src/utils/smartweave"
)

const TypeMsgDataItem = "data_item"

var _ sdk.Msg = &MsgDataItem{}

func (msg *MsgDataItem) Route() string {
	return RouterKey
}

func (msg *MsgDataItem) Type() string {
	return TypeMsgDataItem
}

func (msg *MsgDataItem) GetCreator() sdk.AccAddress {
	pubKey, err := msg.GetPublicKey()
	if err != nil {
		panic(err)
	}
	return sdk.AccAddress(pubKey.Address())
}

func (msg *MsgDataItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.GetCreator()}
}

func (msg *MsgDataItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDataItem) ValidateBasic() (err error) {
	// Data item validation is done in the AnteHandler 
	// to have greater control over it and avoid executing it during recheckTx.
	return nil
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
		Sender:     msg.GetSigners()[0].String(),
	}
}
