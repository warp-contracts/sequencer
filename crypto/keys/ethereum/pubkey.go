package ethereum

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"

	etherum_crypto "github.com/ethereum/go-ethereum/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tmcrypto "github.com/tendermint/tendermint/crypto"
)

type ethereumPK struct {
	public ecdsa.PublicKey
}

func (pk *PubKey) Address() tmcrypto.Address {
	return tmcrypto.AddressHash(pk.Bytes())
}

func (pk *PubKey) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(pk.Address())
}

func (pk *PubKey) VerifySignature(data []byte, signature []byte) bool {
	hashed := sha256.Sum256(data)

	// Get the public key from the signature
	sigPublicKey, err := etherum_crypto.Ecrecover(hashed[:], signature)
	if err != nil {
		return false
	}

	// Check if the public key recovered from the signature matches the owner
	if !bytes.Equal(sigPublicKey, pk.Bytes()) {
		return false
	}
	return true
}

func (pk *PubKey) Bytes() []byte {
	return etherum_crypto.FromECDSAPub(&pk.Key.public)
}

func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	return pk.Type() == other.Type() && bytes.Equal(pk.Bytes(), other.Bytes())
}

func (pk *PubKey) Type() string {
	return "ethereum"
}

func (pk *ethereumPK) Bytes() []byte {
	return etherum_crypto.FromECDSAPub(&pk.public)
}

func (pk *ethereumPK) Size() int {
	return len(pk.Bytes())
}

func (pk *ethereumPK) MarshalTo(data []byte) (int, error) {
	bz := pk.Bytes()
	copy(data, bz)
	return len(bz), nil
}

func (pk *ethereumPK) Unmarshal(bz []byte) error {
	key, err := etherum_crypto.UnmarshalPubkey(bz)
	if err != nil {
		return err
	}
	pk.public = *key
	return nil
}

func UnmarshalPubkey(bz []byte) (*PubKey, error) {
	key, err := etherum_crypto.UnmarshalPubkey(bz)
	if err != nil {
		return nil, err
	}
	return &PubKey{&ethereumPK{*key}}, nil
}

func (pk ethereumPK) MarshalJSON() ([]byte, error) {
	return json.Marshal(pk.public)
}

func (pk *ethereumPK) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &pk.public)
}
