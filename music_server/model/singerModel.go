package model

type Singer struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Index string `xorm:"varchar(5) index" json:"index"` // 索引首字母 	A-Z
	Name string `xorm:"varchar(55)" json:"name"` // 姓名
	Nickname string `xorm:"varchar(255)" json:"nickname"` // 别名
	Avatar string `xorm:"varchar(255)" json:"avatar"` //
	Gender string `xorm:"varchar(10)" json:"gender"` // 性别 男 女
	Area string `xorm:"varchar(20)" json:"area"` // 地区
	Shape string`xorm:"varchar(20)" json:"shape"` // 形式 : 男 女 组合
	Style string `xorm:"varchar(20)" json:"style"` // 风格 : 流行 说唱 国风 摇滚...
}