package model

type Music struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	CategoryId int64 `xorm:"int index" json:"category_id"` // 专辑、歌单 ID
	CategoryTitle string `xorm:"varchar(60)" json:"category_title"` // 专辑名
	SingerId int64 `xorm:"int index" json:"singer_id"` // 歌手ID
	SingerName string `xorm:"varchar(60)" json:"singer_name"` // 歌手名
	Name string `xorm:"varchar(255)" json:"name"` // 歌名
	Link string `xorm:"varchar(255)" json:"link"` // 链接
	Pic string `xorm:"varchar(255)" json:"pic"` // 歌曲配图
}
