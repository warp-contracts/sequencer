package ar

import (
	"github.com/everFinance/goar"
	"github.com/spf13/viper"
	"github.com/warp-contracts/gateway/config"
)

func GetArweaveClient() *goar.Client {
	config.Init()
	url := viper.GetString("arweave.url")
	return goar.NewClient(url)
	//return goar.NewClient("http://localhost:1984")
	//return goar.NewClient("https://arweave.net")
}
