package api

import (
	"cake_mall/serializer"
	userParams "cake_mall/serializer/params/api"
	userService "cake_mall/service/user"
	"cake_mall/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserWeChatLogin 微信登录or注册
func UserWeChatLogin(c *gin.Context) {
	userWeChatLoginParam := &userParams.UserWeChatLoginParam{}
	if err := c.ShouldBind(&userWeChatLoginParam); err == nil {
		res := userService.WeChatLoginService(userWeChatLoginParam.Code, userWeChatLoginParam.EncryptedData, userWeChatLoginParam.Iv)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//BindPhoneAndPwd 仅适用第三方登录后无手机信息，绑定手机和密码
func UserBindPhoneAndPwd(c *gin.Context) {
	userId, _ := c.Get("userId")
	userBindPhoneAndPwdParam := &userParams.UserBindPhoneAndPwdParam{}
	if err := c.ShouldBind(&userBindPhoneAndPwdParam); err == nil {
		res := userService.BindPhoneAndPwdService(userId.(string), userBindPhoneAndPwdParam.Phone, userBindPhoneAndPwdParam.SmsCode, userBindPhoneAndPwdParam.Pwd, util.SmsBindPhoneType)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//UserSmsLogin 验证码登录
func UserSmsLogin(c *gin.Context) {
	userSmsLoginParam := &userParams.UserSmsLoginParam{}
	if err := c.ShouldBind(&userSmsLoginParam); err == nil {
		fmt.Println(userSmsLoginParam)
		res := userService.SmsLogin(userSmsLoginParam.Phone, userSmsLoginParam.SmsCode, util.SmsLoginType)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//UserPwdLogin 密码登录
func UserPwdLogin(c *gin.Context, ) {
	userPwdLoginParam := &userParams.UserPwdLoginParam{}
	if err := c.ShouldBind(&userPwdLoginParam); err == nil {
		res := userService.PwdLogin(userPwdLoginParam.Phone, userPwdLoginParam.Pwd)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//UserRegister 注册
func UserRegister(c *gin.Context, ) {
	userRegisterParam := &userParams.UserRegisterParam{}
	if err := c.ShouldBind(&userRegisterParam); err == nil {
		res := userService.Register(userRegisterParam.Phone, userRegisterParam.Pwd, userRegisterParam.SmsCode, util.SmsRegType)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//UpdatePwd 修改密码
func UserUpdatePwd(c *gin.Context) {
	userUpdatePwdParam := &userParams.UserUpdatePwdParam{}
	if err := c.ShouldBind(&userUpdatePwdParam); err == nil {
		res := userService.UpdatePwdService(userUpdatePwdParam.Phone, userUpdatePwdParam.NewPwd, userUpdatePwdParam.SmsCode, util.SmsUpdatePwdType)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//ResetPwd 重置密码
func UserResetPwd(c *gin.Context) {
	userResetPwdParam := &userParams.UserResetPwdParam{}
	if err := c.ShouldBind(&userResetPwdParam); err == nil {
		res := userService.ResetPwdService(userResetPwdParam.Phone, userResetPwdParam.NewPwd, userResetPwdParam.SmsCode, util.SmsResetPwdType)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}
