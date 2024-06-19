package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var Store *RedisStore
var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "119.45.145.96:6379",
		DB:   0,
	})

	Store = NewRedisStore(rdb, ctx, "captcha_", 10*time.Minute)
	log.Printf("rdb is %v\n", rdb)
}
