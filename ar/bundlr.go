package ar

import (
	"encoding/json"
	"errors"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetBundlr() Bundlr {
	urls := viper.GetStringSlice("arweave.bundlrUrls")
	if urls == nil || len(urls) == 0 {
		logrus.Panic("There are no bundlr URLs in config")
	}
	return &bundlr{
		bundlerUrls: urls,
	}
}

type Bundlr interface {
	UploadToBundlr(transaction *types.Transaction, tags ...types.Tag) (bundlrResp *types.BundlrResp, confirmNode string, err error)
}

type bundlr struct {
	bundlerUrls []string
}

func (b *bundlr) UploadToBundlr(transaction *types.Transaction, tags ...types.Tag) (bundlrResp *types.BundlrResp, confirmNode string, err error) {
	key := viper.GetString("arweave.walletJwk")
	if key == "" {
		err = errors.New("key cannot be empty")
		return
	}
	signer, err := goar.NewSigner([]byte(key))
	if err != nil {
		return
	}
	itemSigner, err := goar.NewItemSigner(signer)
	encodedTransaction, err := json.Marshal(transaction)
	if err != nil {
		return
	}
	item, err := itemSigner.CreateAndSignItem(
		encodedTransaction,
		"", "",
		tags,
	)

	for _, confirmNode = range b.bundlerUrls {
		bundlrResp, err = utils.SubmitItemToBundlr(item, confirmNode)
		if err == nil {
			return
		} else {
			logrus.Warn(err)
		}
	}
	return
}
