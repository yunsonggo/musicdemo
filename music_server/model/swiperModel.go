package model

type Swiper struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Title string `xorm:"varchar(60)" json:"title"`
	Image string `xorm:"varchar(255)" json:"image"`
	Link string `xorm:"varchar(255)" json:"link"`
	RadioId int64 `xorm:"int" json:"radio_id"`
	Sort int `xorm:"bigInt" json:"sort"`
}