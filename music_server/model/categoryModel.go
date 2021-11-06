package model

import "2021/music_demo/music_server/tools"

type Category struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Title string `xorm:"varchar(60)" json:"title"`
	IsTop int `xorm:"int" json:"is_top"`
	Image string `xorm:"varchar(255)" json:"image"`
	Tag string `xorm:"varchar(30)" json:"tag"` // 标签
	Count int `xorm:"int" json:"count"` // 点击次数
	Desc string `xorm:"varchar(255)" json:"desc"` // 简介
	CreatedTime tools.JsonTime `xorm:"created" json:"created_time"`
}
