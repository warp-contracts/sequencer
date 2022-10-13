package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"github.com/warp-contracts/sequencer/ar"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/routes"
)

func main() {
	config.Init()
	ar.StartCacheRead()

	r := gin.New()
	r.Use(ginlogrus.Logger(logrus.StandardLogger()), gin.Recovery())

	r.POST("sequencer/register", routes.RegisterSequencer)
	r.POST("gateway/sequencer/register", routes.RegisterSequencer)

	err := r.Run()
	if err != nil {
		logrus.Panic(err)
	}
}
