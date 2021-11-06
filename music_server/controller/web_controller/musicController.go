package web_controller

import (
	"2021/music_demo/music_server/common"
	"2021/music_demo/music_server/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

var mss = server.NewMusicServer()

// 根据歌手ID 获取歌曲
func GetMusicBySinger(ctx *gin.Context) {
	action := ctx.Query("id")
	id, _ := strconv.ParseInt(action, 10, 64)
	countStr := ctx.Query("count")
	count,_ := strconv.Atoi(countStr)
	startStr := ctx.Query("start")
	start,_ := strconv.Atoi(startStr)
	res, err := mss.MusicFromSinger(id,count,start)
	if err != nil {
		common.Failed(ctx, "获取歌曲数据失败,err = "+err.Error())
		return
	}
	common.Success(ctx, res)
	return
}
