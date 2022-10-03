package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/warp-contracts/gateway/ar"
	"github.com/warp-contracts/gateway/config"
	"github.com/warp-contracts/gateway/routes"
)

func main() {
	config.Init()
	ar.StartCacheRead()

	r := gin.Default()
	r.POST("sequencer/register", routes.RegisterSequencer)

	err := r.Run()
	if err != nil {
		logrus.Panic(err)
	}
}
