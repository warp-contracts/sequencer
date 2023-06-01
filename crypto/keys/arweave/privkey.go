package arweave

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type arweaveSK struct {
	secret rsa.PrivateKey
}

func (sk *PrivKey) Bytes() []byte {
	if sk == nil {
		return nil
	}
	return sk.Key.secret.N.Bytes()
}

func (sk *PrivKey) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	return rsa.SignPSS(rand.Reader, &sk.Key.secret, crypto.SHA256, hashed[:], &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	})
}

func (sk *PrivKey) PubKey() cryptotypes.PubKey {
	return &PubKey{sk.Key.secret.PublicKey.N.Bytes()}
}

func (sk *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	otherPk, ok := other.(*PrivKey)
	if !ok {
		return false
	}
	return sk.Key.secret.Equal(&otherPk.Key.secret)
}

func (sk *PrivKey) Type() string {
	return name
}
