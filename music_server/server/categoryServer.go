package server

import (
	"2021/music_demo/music_server/dial"
	"2021/music_demo/music_server/model"
)

type CategoryServer interface {
	// 获取有限推荐歌单
	Categories(count,start int) (categories []model.Category,err error)
}

type categoryServer struct {}

func NewCategoryServer() CategoryServer {
	return &categoryServer{}
}

// // 获取有限推荐歌单
func (cgs *categoryServer) Categories(count,start int) (categories []model.Category,err error) {
	err = dial.DB.Where("is_top = ?",0).Desc("created_time").Limit(count,start).Find(&categories)
	return
}

