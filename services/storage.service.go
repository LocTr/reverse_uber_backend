package services

import (
	"sync"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

// TODO: Implement the InitMongoDB function without mgm
// func InitMongoDB() {

// 	// setup the mgm default config

var redisDefaultClient *redis.Client
var redisDefaultOnce sync.Once

var redisCache *cache.Cache
var redisCacheOnce sync.Once

func GetRedisDefaultClient() *redis.Client {
	redisDefaultOnce.Do(func() {
		redisDefaultClient = redis.NewClient(&redis.Options{
			Addr: Config.RedisDefaultAddress,
		})
	})
	return redisDefaultClient
}

func GetRedisCache() *cache.Cache {
	redisCacheOnce.Do(func() {
		redisCache = cache.New(&cache.Options{
			Redis:      GetRedisDefaultClient(),
			LocalCache: cache.NewTinyLFU(1000, 0),
		})
	})
	return redisCache
}
