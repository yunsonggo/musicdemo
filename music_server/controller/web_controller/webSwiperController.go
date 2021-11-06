package web_controller

import (
	"2021/music_demo/music_server/common"
	"2021/music_demo/music_server/server"
	"github.com/gin-gonic/gin"
)

var sws = server.NewSwiperServer()
// 获取轮播图数据
func GetSwiperArr(ctx *gin.Context) {
	res,err := sws.SwiperData()
	if err != nil {
		common.Failed(ctx,"获取轮播图数据失败,err = "+err.Error())
		return
	}
	common.Success(ctx,res)
	return
}
