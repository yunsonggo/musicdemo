package server

import (
	"2021/music_demo/music_server/dial"
	"2021/music_demo/music_server/model"
)

type SwiperServer interface {
	// 获取所有轮播数据
	SwiperData() (swiperArr []model.Swiper,err error)
}

type swiperServer struct {}

func NewSwiperServer() SwiperServer {
	return &swiperServer{}
}

// 获取所有轮播数据
func (sws *swiperServer) SwiperData() (swiperArr []model.Swiper,err error) {
	err = dial.DB.Find(&swiperArr)
	return
}