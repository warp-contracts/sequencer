package ar

import (
	"github.com/everFinance/goar"
	"github.com/spf13/viper"
)

func GetArweaveClient() *goar.Client {
	url := viper.GetString("arweave.url")
	conn := goar.NewTempConn()
	conn.SetTempConnUrl(url)
	return conn
}
