package ethereum

import (
	"bytes"
	"crypto/sha256"
	ethereum_crypto "github.com/ethereum/go-ethereum/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	tmcrypto "github.com/tendermint/tendermint/crypto"
)

func (pk *PubKey) Address() tmcrypto.Address {
	return tmcrypto.AddressHash(pk.Bytes())
}

func (pk *PubKey) VerifySignature(data []byte, signature []byte) bool {
	hashed := sha256.Sum256(data)
	return ethereum_crypto.VerifySignature(pk.Key, hashed[:], signature)
}

func (pk *PubKey) Bytes() []byte {
	bz := make([]byte, len(pk.Key))
	copy(bz, pk.Key)
	return bz
}

func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	return pk.Type() == other.Type() && bytes.Equal(pk.Bytes(), other.Bytes())
}

func (pk *PubKey) Type() string {
	return "ethereum"
}

func UnmarshalPubkey(bz []byte) (*PubKey, error) {
	key, err := ethereum_crypto.UnmarshalPubkey(bz)
	if err != nil {
		return nil, err
	}
	publicKeyBytes := ethereum_crypto.FromECDSAPub(key)
	return &PubKey{publicKeyBytes}, nil
}
