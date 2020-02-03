package cache

import (
	"github.com/go-redis/redis/v7"
	"os"
	"time"
)

type Cache struct {
	prefix  string
	redis   *redis.Client //redis连接句柄
	dbIndex int
}

func NewCache(dbIndex int, prefix string) *Cache {
	cache := &Cache{
		prefix:  prefix,
		redis:   newRedis(dbIndex),
		dbIndex: dbIndex,
	}
	return cache
}

func newRedis(dbIndex int) *redis.Client {
	option := &redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       dbIndex,
	}
	client := redis.NewClient(option)
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

func (c *Cache) Set(key string, value interface{}, expireTime time.Duration) *Cache {
	key = c.prefix + key
	c.redis.Set(key, value, expireTime)
	return c
}

func (c *Cache) Get(key string) string {
	key = c.prefix + key
	r := c.redis.Get(key)
	v, _ := r.Result()
	return v
}

func (c *Cache) Delete(key string) int64 {
	key = c.prefix + key
	r := c.redis.Del(key)
	v, _ := r.Result()
	return v
}
