package eth

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/sirupsen/logrus"
)

func Validate(transaction types.Transaction) bool {
	tags, err := utils.TagsDecode(transaction.Tags)
	if err != nil {
		logrus.Warn(err)
		return false
	}
	transaction.Tags = tags
	marshaledTx, err := utils.GetSignatureData(&transaction)
	if err != nil {
		logrus.Warn(err)
		return false
	}
	hash := crypto.Keccak256Hash(marshaledTx)
	signature, err := hexutil.Decode(transaction.Signature)
	if err != nil {
		logrus.Warn(err)
		return false
	}
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		logrus.Warn(err)
		return false
	}
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID)
	return verified
}
