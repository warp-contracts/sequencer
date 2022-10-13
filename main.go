package main

import (
	// gin library
	"github.com/gin-gonic/gin"
	// logrus library
	"github.com/sirupsen/logrus"
	"github.com/warp-contracts/sequencer/ar"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/routes"
)

func main() {
	config.Init()
	ar.StartCacheRead()

	r := gin.Default()
	r.POST("sequencer/register", routes.RegisterSequencer)
	r.POST("gateway/sequencer/register", routes.RegisterSequencer)

	err := r.Run()
	if err != nil {
		logrus.Panic(err)
	}
}
