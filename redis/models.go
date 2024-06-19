package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var V string

type RedisStore struct {
	RedisClient *redis.Client
	Ctx         context.Context
	KeyPrefix   string
	Expiration  time.Duration
}

func NewRedisStore(redisClient *redis.Client, ctx context.Context, keyPrefix string, expiration time.Duration) *RedisStore {
	return &RedisStore{
		RedisClient: redisClient,
		Ctx:         ctx,
		KeyPrefix:   keyPrefix,
		Expiration:  expiration,
	}
}

func (rs *RedisStore) Set(id string, value string) error {
	rs.RedisClient.Set(rs.Ctx, rs.KeyPrefix+id, value, rs.Expiration)
	return nil
}

func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.RedisClient.Get(rs.Ctx, rs.KeyPrefix+id).Result()

	if err != nil {
		log.Printf("get error occurs %v", err)
		return ""
	}
	if clear {
		rs.RedisClient.Del(rs.Ctx, rs.KeyPrefix+id)
	}

	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	V = rs.Get(id, clear)
	return V == answer
}
