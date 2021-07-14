package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuefei7746/go-web-example/utils"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
}

func RedisPoolStats(c *gin.Context) {
	ctx := context.TODO()
	if ping, ok := utils.PingRedis(ctx); !ok {
		renderError(c, fmt.Sprintf("ping redis failed: %s", ping))
	}
	stats := utils.GetRedisCli().PoolStats()
	renderOK(c, stats)
}

func LocalCacheStats(c *gin.Context) {
	stats := utils.GetCacheCli().Stats()
	renderOK(c, stats)
}
