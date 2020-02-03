package main

import (
	"cake_mall/conf"
	"cake_mall/server"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//fmt.Println(util.NowTime())
	//fmt.Println(123123)
	//cos.Auth()
	//s := cos.GetSignature()
	//fmt.Printf("GetSignature:%v\n",s)
	//u2, err := uuid.NewV4()
	//if err != nil {
	//	log.Fatalf("failed to generate UUID: %v", err)
	//}
	//fmt.Printf("generated Version 4 UUID %T", u2)

	//header := req.Header{
	//	"Authorization": cos.Auth(),
	//	"Host":"anyi-cake-12520017765.cos.ap-chengdu.myqcloud.com",
	//}
	//req.Debug = true
	//b, _ := ioutil.ReadFile("1.txt")
	//// 只有url必选，其它参数都是可选
	//r, err := req.Put("http://anyi-cake-12520017765.cos.ap-chengdu.myqcloud.com", header,b)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%+v", r) // 打印详细信息
	conf.Init()
	r := server.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}