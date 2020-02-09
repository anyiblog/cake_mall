package main

import (
	"cake_mall/conf"
	"cake_mall/server"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conf.Init()
	r := server.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
