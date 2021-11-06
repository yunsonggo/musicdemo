package server

import (
	"2021/music_demo/music_server/dial"
	"2021/music_demo/music_server/model"
)

type MusicServer interface {
	// 根据歌手ID 查询 歌曲
	MusicFromSinger(id int64,count,start int) (musicList []model.Music,err error)
}

type musicServer struct {}

func NewMusicServer() MusicServer {
	return &musicServer{}
}

// 根据歌手ID 查询 歌曲
func (mss *musicServer) MusicFromSinger(id int64,count,start int) (musicList []model.Music,err error) {
	err = dial.DB.Where("singer_id = ?",id).Limit(count,start).Find(&musicList)
	return
}