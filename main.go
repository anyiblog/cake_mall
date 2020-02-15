package main

import (
	"cake_mall/conf"
	"cake_mall/server"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
)

// 登录拦截中间件
func main() {
	conf.Init()
	r := server.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
