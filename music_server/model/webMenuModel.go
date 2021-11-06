package model

type WebMenu struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Title string `xorm:"varchar(30)" json:"title"`
	Path string `xorm:"varchar(255)" json:"path"`
	Sort int `xorm:"int" json:"sort"`
}
