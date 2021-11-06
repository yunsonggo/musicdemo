package web_controller

import (
	"2021/music_demo/music_server/common"
	"2021/music_demo/music_server/server"
	"github.com/gin-gonic/gin"
)

var sgs = server.NewSingerServer()

// 所有歌手列表
func GetSingerList(ctx *gin.Context) {
	res, err := sgs.SingerListInfo()
	if err != nil {
		common.Failed(ctx,"获取歌手列表失败,err= " + err.Error())
		return
	}
	common.Success(ctx,res)
	return
}
