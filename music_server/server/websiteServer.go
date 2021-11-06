package server

import (
	"2021/music_demo/music_server/dial"
	"2021/music_demo/music_server/model"
)

type WebsiteServer interface {
	// 网站信息
	WebsiteInfo() (website model.Website,err error)
}

type websiteServer struct {}

func NewWebsiteServer() WebsiteServer {
	return &websiteServer{}
}

// 网站信息
func (wss *websiteServer) WebsiteInfo() (website model.Website,err error) {
	_,err = dial.DB.Where("id = 1").Get(&website)
	return
}