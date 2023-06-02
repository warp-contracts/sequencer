package arweave

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"math/big"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tmcrypto "github.com/tendermint/tendermint/crypto"
)

func (pk *PubKey) Address() tmcrypto.Address {
	return tmcrypto.AddressHash(pk.Key)
}

func (pk *PubKey) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(pk.Address())
}

func (pk *PubKey) VerifySignature(data []byte, signature []byte) bool {
	hashed := sha256.Sum256(data)

	ownerPublicKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes([]byte(pk.Key)),
		E: 65537, //"AQAB"
	}

	return rsa.VerifyPSS(ownerPublicKey, crypto.SHA256, hashed[:], []byte(signature), &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	}) != nil
}

func (pk *PubKey) Bytes() []byte {
	return pk.Key
}

func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	return pk.Type() == other.Type() && bytes.Equal(pk.Bytes(), other.Bytes())
}

func (pk *PubKey) Type() string {
	return name
}
