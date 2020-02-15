package conf

import (
	"github.com/go-redis/redis/v7"
	"os"
)

// RedisConnection 服务启动时测试Redis是否就绪
func RedisConnection(dbIndex int) {
	option := &redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       dbIndex,
	}
	client := redis.NewClient(option)
	_, err := client.Ping().Result()
	if err != nil {
		panic("连接Redis失败：" + err.Error())
	}
	// 仅仅启动服务时，测试Redis连接。
	_ = client.Close()
}
