package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/warp-contracts/sequencer/ar"
)

func ArweaveInfoRoute(r *gin.Context) {
	r.JSON(200, ar.GetCachedInfo().NetworkInfo)
}
