# 风铃甜点屋服务端

## 引用项目

本项目引用的三方包如下：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
4. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
5. 腾讯云Cos，腾讯云短信
6. 本项目是使用基于redis缓存实现的token来保存登录状态

## 接口规划

1. server/router目录里/v1是小程序用户端接口
2. /admin是服务端接口

## 模块规划

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api和admin文件夹就是MVC框架的controller，负责接参传参
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理业务逻辑
4. serializer储存通用的json模型
5. cache负责redis缓存相关的代码
6. middleware Gin中间件
7. tmp 临时缓存文件
8. util一些通用的工具包
9. conf 程序基本配置
10. .env 配置文件

## 运行说明

项目启动依赖mysql和redis

可以参照.example-env填写腾讯云相关配置

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)进行依赖管理


## 运行

```shell
## 运行
go run main.go

## 编译（参考Makefile文件）
make 平台名称 （make mac）

```

项目运行后启动在8080端口（可以修改，参考gin文档)