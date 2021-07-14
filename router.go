package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuefei7746/go-web-example/controller"
	"github.com/yuefei7746/go-web-example/utils"
)

func initRouter() *gin.Engine {
	utils.Logger.Debug("Initial router")

	router := gin.Default()

	router.GET("/ping", controller.Ping)

	info := router.Group("/handler", gin.BasicAuth(controller.AuthorizationAccount))
	{
		info.GET("/redis_stats", controller.RedisPoolStats)
		info.GET("/cache_stats", controller.LocalCacheStats)

		pprofRouter := info.Group("/pprof")
		{
			pprofRouter.GET("/", controller.PprofIndex)
			pprofRouter.GET("/cmdline", controller.PprofCmdline)
			pprofRouter.GET("/profile", controller.PprofProfile)
			pprofRouter.GET("/symbol", controller.PprofSymbol)
			pprofRouter.GET("/trace", controller.PprofTrace)
		}
	}

	return router
}
