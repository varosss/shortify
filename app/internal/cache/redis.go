package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

const URL_CACHE_DURATION = 2 * time.Hour

func NewRedisClient(redisAddr string) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	return redisClient
}
