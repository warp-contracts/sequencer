package ethereum

import (
	"bytes"
	"crypto/sha256"

	ethereum_crypto "github.com/ethereum/go-ethereum/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	tmcrypto "github.com/tendermint/tendermint/crypto"
)

func (pk *PubKey) Address() tmcrypto.Address {
	return tmcrypto.AddressHash(pk.Key)
}

func (pk *PubKey) VerifySignature(data []byte, signature []byte) bool {
	if len(signature) == ethereum_crypto.SignatureLength {
		// remove recovery ID (V) if contained in the signature
		signature = signature[:len(signature)-1]
	}

	hashed := sha256.Sum256(data)
	return ethereum_crypto.VerifySignature(pk.Key, hashed[:], signature)
}

func (pk *PubKey) Bytes() []byte {
	return pk.Key
}

func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	return pk.Type() == other.Type() && bytes.Equal(pk.Bytes(), other.Bytes())
}

func (pk *PubKey) Type() string {
	return "ethereum"
}

func FromOwner(owner []byte) (*PubKey, error) {
	key, err := ethereum_crypto.UnmarshalPubkey(owner)
	if err != nil {
		return nil, err
	}
	publicKeyBytes := ethereum_crypto.FromECDSAPub(key)
	return &PubKey{publicKeyBytes}, nil
}
