package model

type Website struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(30)" json:"name"`	// 网站名称
	Logo string `xorm:"varchar(255)" json:"logo"`	// logo地址
	Static string `xorm:"varchar(60)" json:"static"` // 静态文件访问前缀
	Url string `xorm:"varchar(255)" json:"url"`		// 域名
}
