package ar

import (
	"github.com/everFinance/goar/types"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/warp-contracts/gateway/config"
	"time"
)

var cacheTaskStarted = false
var cachedInfo *types.NetworkInfo

func StartCacheRead() {
	config.Init("../")
	if cacheTaskStarted {
		return
	}
	cacheInitedChannel := make(chan bool)
	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(30).Millisecond().Do(func() {
		defer func() { cacheInitedChannel <- true }()
		arweaveClient := GetArweaveClient()
		info, err := arweaveClient.GetInfo()
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Debug(info)
			cachedInfo = info
		}
	})
	if err != nil {
		logrus.Panic(err)
	}
	scheduler.StartAsync()
	<-cacheInitedChannel
	cacheTaskStarted = true
}

func GetCachedInfo() types.NetworkInfo {
	if !cacheTaskStarted {
		logrus.Panic("Read cache task is not started")
	}

	return *cachedInfo
}
