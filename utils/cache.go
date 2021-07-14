package utils

import (
	"github.com/bytedance/sonic"
	"github.com/go-redis/cache/v8"
	"sync"
	"time"
)

var (
	cacheClient *cache.Cache
	cacheOnce   sync.Once
)

func GetCacheCli() *cache.Cache {
	cacheOnce.Do(initCacheCli)
	return cacheClient
}

func initCacheCli() {
	cacheClient = cache.New(&cache.Options{
		//Redis:        GetRedisCli(),
		LocalCache:   cache.NewTinyLFU(10, time.Minute),
		StatsEnabled: true,
		Marshal:      sonic.Marshal,
		Unmarshal:    sonic.Unmarshal,
	})
	Logger.Info("memory cache ready...")
}
