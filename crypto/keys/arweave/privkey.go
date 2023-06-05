package arweave

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type arweaveSK struct {
	secret rsa.PrivateKey
}

func (sk *PrivKey) Bytes() []byte {
	return sk.Key.Bytes()
}

func (sk *PrivKey) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	return rsa.SignPSS(rand.Reader, &sk.Key.secret, crypto.SHA256, hashed[:], &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	})
}

func (sk *PrivKey) PubKey() cryptotypes.PubKey {
	return &PubKey{&arweavePK{sk.Key.secret.PublicKey}}
}

func (sk *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	otherSk, ok := other.(*PrivKey)
	if !ok {
		return false
	}
	return sk.Key.secret.Equal(&otherSk.Key.secret)
}

func (sk *PrivKey) Type() string {
	return "arweave"
}

func (sk *arweaveSK) Bytes() []byte {
	return x509.MarshalPKCS1PrivateKey(&sk.secret)
}

func (sk *arweaveSK) MarshalTo(data []byte) (int, error) {
	bz := sk.Bytes()
	copy(data, bz)
	return len(bz), nil
}

func (sk *arweaveSK) Unmarshal(bz []byte) error {
	key, err := x509.ParsePKCS1PrivateKey(bz)
	if err != nil {
		return err
	}
	sk.secret = *key
	return nil
}

func (sk *arweaveSK) Size() int {
	return len(sk.Bytes())
}

