package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	Set(key string, value string, expirationTime time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

type RedisCache struct {
	Client *redis.Client
}

func (cache *RedisCache) Set(key string, value string, expirationTime time.Duration) error {
	return cache.Client.Set(key, value, expirationTime).Err()
}

func (cache *RedisCache) Get(key string) (string, error) {
	return cache.Client.Get(key).Result()
}

func (cache *RedisCache) Delete(key string) error {
	return cache.Client.Del(key).Err()
}
