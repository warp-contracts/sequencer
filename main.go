package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginlogrus "github.com/toorop/gin-logrus"
	"github.com/warp-contracts/sequencer/ar"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/db/conn"
	"github.com/warp-contracts/sequencer/routes"
)

func main() {
	config.Init()
	ar.StartCacheRead()

	basicChecks()

	gin.SetMode(viper.GetString("sequencer.mode"))
	r := gin.New()
	r.Use(ginlogrus.Logger(logrus.StandardLogger()), gin.Recovery())

	//r.POST("sequencer/register", routes.RegisterSequencer)
	r.POST("gateway/sequencer/register", routes.RegisterSequencer)
	r.GET("gateway/arweave/info", routes.ArweaveInfoRoute)
	r.GET("gateway/arweave/block", routes.ArweaveBlockRoute)

	err := r.Run(":" + viper.GetString("sequencer.port"))
	if err != nil {
		logrus.Panic(err)
	}
}

func basicChecks() {
	conn.GetConnection()
	ar.GetBundlr()
	if ar.GetCachedInfo() == nil {
		logrus.Panic("Cached information is empty")
	}
}
