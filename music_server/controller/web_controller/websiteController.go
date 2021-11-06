package web_controller

import (
	"2021/music_demo/music_server/common"
	"2021/music_demo/music_server/server"
	"github.com/gin-gonic/gin"
)

var websiteServer = server.NewWebsiteServer()

// 获取网站信息
func GetWebsiteInfo(ctx *gin.Context) {
	website,err := websiteServer.WebsiteInfo()
	if err != nil {
		common.Failed(ctx,"获取网站信息失败,err = " + err.Error())
		return
	}
	common.Success(ctx,website)
	return
}