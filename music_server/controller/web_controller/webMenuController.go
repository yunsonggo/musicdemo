package web_controller

import (
	"2021/music_demo/music_server/common"
	"2021/music_demo/music_server/server"
	"github.com/gin-gonic/gin"
)

var wms = server.NewWebMenuServer()
// 网站前端菜单
func GetWebMenus(ctx *gin.Context) {
	webMenus,err := wms.WebMenus()
	if err != nil {
		common.Failed(ctx,"获取菜单数据失败,err = " + err.Error())
		return
	}
	common.Success(ctx,webMenus)
	return
}