package ar

import (
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/warp-contracts/sequencer/config"
	"time"
)

var cacheTaskStarted = false
var blockInfo *BlockInfo

func StartCacheRead() {
	config.Init()
	if cacheTaskStarted {
		return
	}
	cacheInitedChannel := make(chan bool)
	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(30).Seconds().Do(func() {
		defer func() { cacheInitedChannel <- true }()
		arweaveClient := GetArweaveClient()
		info, err := arweaveClient.GetInfo()
		if err != nil {
			logrus.Error("Error with reading Arweave network info", err)
		} else {
			logrus.Debug(info)
			block, err := arweaveClient.GetBlockByID(info.Current)
			if err != nil {
				logrus.Error(err)
			} else {
				blockInfo = &BlockInfo{
					NetworkInfo:  info,
					CurrentBlock: block,
				}
			}
		}
	})
	if err != nil {
		logrus.Panic(err)
	}
	scheduler.StartAsync()
	<-cacheInitedChannel
	cacheTaskStarted = true
}

func GetCachedInfo() *BlockInfo {
	if !cacheTaskStarted {
		logrus.Panic("Read cache task is not started")
	}
	return blockInfo
}
