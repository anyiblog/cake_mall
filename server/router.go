package server

import (
	"cake_mall/admin"
	"cake_mall/api"
	"cake_mall/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.Cors())

	{
		//根接口
		r.GET("SendSms", api.SendSms)
		r.POST("Ping", api.Ping)
		r.POST("UploadFile", api.UploadFile)
		r.POST("DeleteFile", api.DeleteFile)
	}

	c := r.Group("/v1") //用户端接口 client
	{
		c.POST("PwdLogin", api.UserPwdLogin)       //密码登录
		c.POST("WeChatLogin", api.UserWeChatLogin) //微信登录
		c.POST("SmsLogin", api.UserSmsLogin)       //验证码登录
		c.POST("Register", api.UserRegister)       //注册
		c.POST("ResetPwd", api.UserResetPwd)       //重置密码

		cUser := c.Group("/User")
		cUser.Use(middleware.TokenAuth()) //Token验证中间件
		{
			cUser.POST("BindPhoneAndPwd", api.UserBindPhoneAndPwd) //第三方登录后，设置手机密码
			cUser.POST("UpdatePwd", api.UserUpdatePwd)             //更新密码（找回密码）
		}
		cHome := c.Group("/Home")
		cUser.Use(middleware.TokenAuth()) //Token验证中间件
		{
			cHome.GET("Banner", api.HomeBanner)        //Banner图标
			cHome.GET("Nav", api.HomeBanner)           //金刚区图标
			cHome.POST("UpdatePwd", api.UserUpdatePwd) //更新密码（找回密码）
		}
	}

	s := r.Group("/admin") //后台端接口 server
	{
		s.POST("Login", admin.UserPwdLogin)
		s.POST("UserInfo", admin.UserInfo)

		s.Use(middleware.TokenAuth())
		s.Use(middleware.AdminAuth()) //是否管理员
		{
			s.GET("CakeList", admin.GetCakeList)
			s.GET("TagList", admin.GetTagList)
			s.GET("ImgList", admin.GetImgListForTag)
		}

		// 需要登录保护的（验证Token）
		sUser := s.Group("/User")
		sUser.Use(middleware.TokenAuth())
		{
			sUser.POST("Logout", admin.Logout)
			sUser.POST("Ping", api.Ping)
		}
	}
	return r
}
