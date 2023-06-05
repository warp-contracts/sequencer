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

type arweavePK struct {
	public rsa.PublicKey
}

func (pk *PubKey) Address() tmcrypto.Address {
	return tmcrypto.AddressHash(pk.Bytes())
}

func (pk *PubKey) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(pk.Address())
}

func (pk *PubKey) VerifySignature(data []byte, signature []byte) bool {
	hashed := sha256.Sum256(data)

	return rsa.VerifyPSS(&pk.Key.public, crypto.SHA256, hashed[:], []byte(signature), &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	}) != nil
}

func (pk *PubKey) Bytes() []byte {
	return pk.Key.Bytes()
}

func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	return pk.Type() == other.Type() && bytes.Equal(pk.Bytes(), other.Bytes())
}

func (pk *PubKey) Type() string {
	return "arweave"
}

func (pk *arweavePK) Bytes() []byte {
	return pk.public.N.Bytes()
}

func (pk *arweavePK) Size() int {
	return len(pk.Bytes())
}

func (pk *arweavePK) MarshalTo(data []byte) (int, error) {
	bz := pk.Bytes()
	copy(data, bz)
	return len(bz), nil
}

func unmarshalRsaPublicKey(bz []byte) rsa.PublicKey {
	return rsa.PublicKey{
		N: new(big.Int).SetBytes(bz),
		E: 65537, //"AQAB"
	}
} 

func (pk *arweavePK) Unmarshal(bz []byte) error {
	pk.public = unmarshalRsaPublicKey(bz)
	return nil
}

func UnmarshalPubkey(bz []byte) *PubKey {
	return &PubKey{&arweavePK{unmarshalRsaPublicKey(bz)}}
}
