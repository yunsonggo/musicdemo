package web_controller

import (
	"2021/music_demo/music_server/common"
	"2021/music_demo/music_server/param"
	"2021/music_demo/music_server/server"
	"github.com/gin-gonic/gin"
)

var cgs = server.NewCategoryServer()
// // 获取有限推荐歌单
func GetTopLimitCategory(ctx *gin.Context) {
	var countStartParam param.CountStartParam
	err := ctx.ShouldBind(&countStartParam)
	if err != nil {
		common.Failed(ctx,"获取歌单参数失败,err = " + err.Error())
		return
	}
	res,err := cgs.Categories(countStartParam.Count,countStartParam.Start)
	if err != nil {
		common.Failed(ctx,"获取歌单数据失败,err = " + err.Error())
		return
	}
	common.Success(ctx,res)
	return
}
