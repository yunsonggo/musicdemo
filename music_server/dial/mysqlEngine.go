package dial

import (
	"2021/music_demo/music_server/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

func MysqlEngine(conn,driver string , show bool) {
	db,err := xorm.NewEngine(driver,conn)
	if err != nil {
		panic(fmt.Sprintf("连接数据库异常:%v\n",err))
	}
	db.ShowSQL(show)
	//err = db.Sync2(
	//	new(model.Website),
	//	)
	err = createTable(db)
	if err != nil {
		panic(fmt.Sprintf("创建数据表出现异常:%v\n",err))
	}
	DB = db
	return
}

// 创建数据表
func createTable(db *xorm.Engine) error {
	err := db.Sync2(
		new(model.Website),
		new(model.WebMenu),
		new(model.Swiper),
		new(model.Category),
		new(model.Singer),
		new(model.Music),
	)
	return err
}