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
	UploadToBundlr(transaction *types.Transaction, tags ...types.Tag) (*types.BundlrResp, error)
}

type bundlr struct {
	bundlerUrls []string
}

func (b *bundlr) UploadToBundlr(transaction *types.Transaction, tags ...types.Tag) (*types.BundlrResp, error) {
	var err error
	key := viper.GetString("arweave.walletJwk")
	if key == "" {
		return nil, errors.New("key cannot be empty")
	}
	signer, err := goar.NewSigner([]byte(key))
	if err != nil {
		return nil, err
	}
	itemSigner, err := goar.NewItemSigner(signer)
	encodedTransaction, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}
	item, err := itemSigner.CreateAndSignItem(
		encodedTransaction,
		"", "",
		tags,
	)

	for _, url := range b.bundlerUrls {
		var resp *types.BundlrResp
		resp, err = utils.SubmitItemToBundlr(item, url)
		if err == nil {
			return resp, nil
		} else {
			logrus.Warn(err)
		}
	}
	return nil, err
}
