package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/warp-contracts/sequencer/crypto/keys/arweave"
	"github.com/warp-contracts/sequencer/crypto/keys/ethereum"

	"github.com/warp-contracts/syncer/src/utils/bundlr"
)

func BroadcastDataItem(ctx client.Context, dataItem MsgDataItem) (*sdk.TxResponse, error) {
	tx, err := createTxWithDataItem(ctx, dataItem)
	if err != nil {
		return nil, err
	}

	txBytes, err := ctx.TxConfig.TxEncoder()(tx)
	if err != nil {
		return nil, err
	}

	// Validates the message and sends it out
	return ctx.BroadcastTx(txBytes)
}

func createTxWithDataItem(ctx client.Context, dataItem MsgDataItem) (tx authsigning.Tx, err error) {
	txBuilder := ctx.TxConfig.NewTxBuilder()

	err = txBuilder.SetMsgs(&dataItem)
	if err != nil {
		return
	}

	signature, err := getSignature(dataItem)
	if err != nil {
		return
	}
	err = txBuilder.SetSignatures(signature)
	if err != nil {
		return
	}

	tx = txBuilder.GetTx()
	return
}

func getSignature(dataItem MsgDataItem) (signature txsigning.SignatureV2, err error) {
	pubKey, err := getPublicKey(dataItem)
	if err != nil {
		return
	}
	sequence, err := dataItem.GetSequenceFromTags()
	if err != nil {
		return
	}
	signature = txsigning.SignatureV2{
		PubKey:   pubKey,
		Sequence: sequence,
		Data:     nil,
	}
	return
}

func getPublicKey(dataItem MsgDataItem) (cryptotypes.PubKey, error) {
	return GetPublicKey(dataItem.DataItem.SignatureType, dataItem.DataItem.Owner)
}

func GetPublicKey(signatureType bundlr.SignatureType, owner []byte) (cryptotypes.PubKey, error) {
	switch signatureType {
	case bundlr.SignatureTypeArweave:
		pubKey := arweave.UnmarshalPubkey(owner)
		return pubKey, nil
	case bundlr.SignatureTypeEtherum:
		pubKey, err := ethereum.UnmarshalPubkey(owner)
		return pubKey, err
	default:
		return nil, bundlr.ErrUnsupportedSignatureType
	}
}
