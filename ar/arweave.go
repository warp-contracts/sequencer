package ar

import (
	"github.com/everFinance/goar"
	"github.com/spf13/viper"
)

func GetArweaveClient() *goar.Client {
	return goar.NewClient(viper.GetString("arweave.url"))
}
