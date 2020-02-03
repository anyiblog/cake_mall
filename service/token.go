package service

import (
	"cake_mall/cache"
	"cake_mall/serializer"
	"cake_mall/util"
	"time"
)

// 检测Redis里面的Token是否跟数据库一致
func CheckTokenRedis() {

}

/*
130用户

A 登录 130 后 mysql token 1
redis token 1

B 登录后 130后 mysql token 2

 检查redis 时候有对应的Token， redis token 2



*/

//保存token到redis
func SaveTokenRedis(userId, token string) {
	c := cache.NewCache(util.RedisToken, "token_")
	key := userId
	value := token
	c.Set(key, value, serializer.TokenRedisTime*(24*time.Hour))
}

func DeleteTokenRedis(userId string) bool {
	c := cache.NewCache(util.RedisToken, "token_")
	key := userId
	if c.Delete(key) > 0 {
		return true
	} else {
		return false
	}
}
