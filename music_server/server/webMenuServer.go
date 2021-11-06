package server

import (
	"2021/music_demo/music_server/dial"
	"2021/music_demo/music_server/model"
)

type WebMenuServer interface {
	// 获取前端菜单
	WebMenus() (webMenus []model.WebMenu,err error)
}

type webMenuServer struct {}

func NewWebMenuServer() WebMenuServer {
	return &webMenuServer{}
}

// 获取前端菜单
func (wms *webMenuServer) WebMenus() (webMenus []model.WebMenu,err error) {
	err = dial.DB.Find(&webMenus)
	return
}