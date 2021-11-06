package router

import (
	"2021/music_demo/music_server/common"
	"2021/music_demo/music_server/config"
	"2021/music_demo/music_server/controller/backend_controller"
	"2021/music_demo/music_server/controller/web_controller"
	"2021/music_demo/music_server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/yakaa/log4g"
	"net/http"
)

func NewRouter() *gin.Engine {
	wc := config.Conf.Website
	// 判断是否启用日志
	if wc.WebsiteMode == gin.ReleaseMode {
		log4g.Init(log4g.Config{Path: wc.WebsiteLog,Stdout: true})
		gin.DefaultWriter = log4g.InfoLog
		gin.DefaultErrorWriter = log4g.ErrorLog
	}
	// 初始化引擎
	app := gin.Default()
	// 设置静态文件路由
	app.StaticFS("/api/static",http.Dir("./public"))
	// 启用session
	common.InitSession(app)
	// 启用cors
	app.Use(middleware.CorsMiddleware())
	// web 接口
	// 自由访问
	wg := app.Group("/api")
	{
		// 获取网站信息
		wg.GET("/website/info",web_controller.GetWebsiteInfo)
		// 获取网站菜单
		wg.GET("/menus",web_controller.GetWebMenus)
		// 获取轮播数据
		wg.GET("/swiper",web_controller.GetSwiperArr)
		// 获取歌单数据
		wg.POST("/category",web_controller.GetTopLimitCategory)
		// 获取歌手列表
		wg.GET("/singer",web_controller.GetSingerList)
		// 根据参数获取歌曲
		wg.GET("/music/singer",web_controller.GetMusicBySinger)
		// 需要登陆权限
		wag := wg.Group("/auth")
		wag.Use(middleware.WebAuth())
		{
			wag.GET("/ping",web_controller.GetAuthPing)
		}
	}
	// 后台接口
	// 自由访问
	bg := app.Group("/backend")
	{
		bg.GET("/ping",backend_controller.GetPing)
		// 后台需要权限
		bag := bg.Group("/auth")
		bag.Use(middleware.BackendAuth())
		{
			bag.GET("/ping",backend_controller.GetAuthPing)
		}
	}
	return app
}