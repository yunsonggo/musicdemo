package config

import "2021/music_demo/music_server/dial"

func InitEngine() {
	mc := Conf.Mysql
	rc := Conf.Redis
	dial.MysqlEngine(mc.MysqlAddr,mc.MysqlDriver,mc.MysqlShow)
	dial.InitRedisClient(rc.RedisAddr,rc.RedisPwd)
}