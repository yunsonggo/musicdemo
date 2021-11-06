package server

import (
	"2021/music_demo/music_server/dial"
	"2021/music_demo/music_server/model"
	"fmt"
)

type SingerServer interface {
	// 获取所有歌手索引数据
	SingerListInfo() (singerListInfo []SingerList,err error)
}

type singerServer struct {}

type SingerList struct {
	Index string `json:"index"`
	Singers []model.Singer `json:"singers"`
}

func NewSingerServer() SingerServer {
	return &singerServer{}
}

// 获取所有歌手数据
func (sgs *singerServer) singers() (singers []model.Singer,err error) {
	err = dial.DB.Find(&singers)
	return
}

// 获取所有歌手索引数据
func (sgs *singerServer) SingerListInfo() (singerListInfo []SingerList,err error) {
	res,err := sgs.singers()
	if err != nil {
		return
	}
	for index := 'A'; index <= 'Z'; index ++ {
		singerList := new(SingerList)
		singerList.Index = fmt.Sprintf("%c",index)
		for _,singer := range res {
			if singer.Index == singerList.Index {
				singerList.Singers = append(singerList.Singers, singer)
			}
		}
		singerListInfo = append(singerListInfo, *singerList)
	}
	return
}