package crypt

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/warp-contracts/sequencer/crypt/p256/secp256k1"
	"github.com/warp-contracts/sequencer/crypt/p256/secp256k1/vrf"
	"github.com/warp-contracts/sequencer/crypt/p256/secp256k1/vrf/secp256k1VRF"
	"math/big"
	"strings"
)

var (
	curve = secp256k1.S256()
	key   vrf.PrivateKey
)

func GetKey() vrf.PrivateKey {
	if key == nil {
		k, err := UnmarshalKey(viper.GetString("vrf.privateKey"))
		if err != nil {
			logrus.Panic(err)
		}
		key = k
	}
	return key
}

func MarshalKey(key vrf.PrivateKey) string {
	privateKey := key.(*secp256k1VRF.PrivateKey)
	dst := make([]byte, len(privateKey.D.Bytes())*2)
	hex.Encode(dst, privateKey.D.Bytes())
	return string(dst)
}

func UnmarshalKey(key string) (vrf.PrivateKey, error) {
	if key == "" {
		return nil, errors.New("key cannot be empty")
	}

	k, err := hex.DecodeString(strings.TrimSpace(key))
	if err != nil {
		return nil, err
	}

	ecdsaPriv, err := ToECDSAPrivateKey(k)
	if err != nil {
		fmt.Printf("ecdsa err: %v", err)
	}
	return secp256k1VRF.NewVRFSigner(ecdsaPriv)
}

// ToECDSAPrivateKey creates a private key with the given data value.
func ToECDSAPrivateKey(d []byte) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = new(big.Int).SetBytes(d)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	return priv, nil
}

// GetCompactPublicKey
// The implementation was taken from:
// https://github.com/indutny/elliptic/blob/43ac7f230069bd1575e1e4a58394a512303ba803/dist/elliptic.js#L314
func GetCompactPublicKey(k vrf.PrivateKey) string {
	pub := k.Public().(*ecdsa.PublicKey)
	var y []byte
	if isEven(pub.Y) {
		y = []byte{0x02}
	} else {
		y = []byte{0x03}
	}
	return hex.EncodeToString(append(y, pub.X.Bytes()...))
}

func isEven(i *big.Int) bool {
	return i.Bit(0) == 0
}
