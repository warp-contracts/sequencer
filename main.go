package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warp-contracts/gateway/routes"
)

func main() {
	r := gin.Default()
	r.POST("sequencer/register", routes.RegisterSequencer)
}
