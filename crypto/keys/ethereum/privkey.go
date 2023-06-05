package ethereum

import (
	"crypto/ecdsa"
	"crypto/sha256"
	etherum_crypto "github.com/ethereum/go-ethereum/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/warp-contracts/syncer/src/utils/bundlr"
)

type ethereumSK struct {
	secret ecdsa.PrivateKey
}

func (sk *PrivKey) Bytes() []byte {
	return sk.Key.Bytes()
}

func (sk *PrivKey) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	return etherum_crypto.Sign(hashed[:], &sk.Key.secret)
}

func (sk *PrivKey) PubKey() cryptotypes.PubKey {
	publicKeyECDSA, ok := sk.Key.secret.Public().(*ecdsa.PublicKey)
	if !ok {
		panic(bundlr.ErrFailedToParseEtherumPublicKey)
	}
	return &PubKey{&ethereumPK{*publicKeyECDSA}}
}

func (sk *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	otherSk, ok := other.(*PrivKey)
	if !ok {
		return false
	}
	return sk.Key.secret.Equal(&otherSk.Key.secret)
}

func (sk *PrivKey) Type() string {
	return "ethereum"
}

func (sk *ethereumSK) Bytes() []byte {
	return etherum_crypto.FromECDSA(&sk.secret)
}

func (sk *ethereumSK) MarshalTo(data []byte) (int, error) {
	bz := sk.Bytes()
	copy(data, bz)
	return len(bz), nil
}

func (sk *ethereumSK) Unmarshal(bz []byte) (err error) {
	key, err := etherum_crypto.ToECDSA(bz)
	if err != nil {
		return err
	}
	sk.secret = *key
	return nil
}

func (sk *ethereumSK) Size() int {
	return len(sk.Bytes())
}
