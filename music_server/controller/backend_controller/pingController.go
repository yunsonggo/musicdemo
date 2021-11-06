package backend_controller

import (
	"2021/music_demo/music_server/common"
	"github.com/gin-gonic/gin"
)

func GetPing(ctx *gin.Context) {
	common.Success(ctx,"pong")
	return
}

func GetAuthPing(ctx *gin.Context) {
	common.Success(ctx,"auth pong")
	return
}