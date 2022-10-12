package p256

import (
	"errors"
	"github.com/warp-contracts/sequencer/crypt/p256/keystore"
	"github.com/warp-contracts/sequencer/crypt/p256/secp256k1"
)

var (
	// ErrAlgorithmInvalid invalid Algorithm for sign.
	ErrAlgorithmInvalid = errors.New("invalid algorithm")
)

func NewPrivateKey(alg keystore.Algorithm, data []byte) (keystore.PrivateKey, error) {
	switch alg {
	case keystore.SECP256K1:
		var (
			priv *secp256k1.PrivateKey
			err  error
		)
		if len(data) == 0 {
			priv = secp256k1.GeneratePrivateKey()
		} else {
			priv = new(secp256k1.PrivateKey)
			err = priv.Decode(data)
		}
		if err != nil {
			return nil, err
		}
		return priv, nil
	default:
		return nil, ErrAlgorithmInvalid
	}
}
