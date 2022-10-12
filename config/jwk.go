package config

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var jwkKey jwk.Key

func GetArConnectAsJwkKey() jwk.Key {
	if jwkKey == nil {
		str := viper.GetString("arweave.arConnectKey")
		k, err := jwk.ParseKey([]byte(str))
		if err != nil {
			logrus.Panic(err)
		}
		jwkKey = k
	}
	return jwkKey
}
