package db_redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func InitDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}
