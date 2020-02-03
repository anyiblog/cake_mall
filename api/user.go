package api

import (
	userService "cake_mall/service/user"
	"cake_mall/util"
	"github.com/gin-gonic/gin"
)

// UserWeChatLogin 微信登录or注册
func UserWeChatLogin(c *gin.Context) {
	jsCode := c.Query("code")
	encryptedData := c.Query("encryptedData")
	iv := c.Query("iv")
	res := userService.WeChatLoginService(jsCode, encryptedData, iv)
	c.JSON(200, res)
}

//BindPhoneAndPwd 仅适用第三方登录后无手机信息，绑定手机和密码
func UserBindPhoneAndPwd(c *gin.Context) {
	userId, _ := c.Get("userId")
	phone := c.PostForm("phone")
	smsCode := c.PostForm("smsCode")
	pwd := c.PostForm("pwd")
	res := userService.BindPhoneAndPwdService(userId.(string), phone, smsCode, pwd, util.SmsBindPhoneType)
	c.JSON(200, res)
}

//UserSmsLogin 验证码登录
func UserSmsLogin(c *gin.Context, ) {
	phone := c.PostForm("phone")
	smsCode := c.PostForm("smsCode")
	res := userService.SmsLogin(phone, smsCode, util.SmsLoginType)
	c.JSON(200, res)
}

//UserPwdLogin 密码登录
func UserPwdLogin(c *gin.Context, ) {
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")
	res := userService.PwdLogin(phone, pwd)
	c.JSON(200, res)
}

//UserRegister 注册
func UserRegister(c *gin.Context, ) {
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")
	smsCode := c.PostForm("smsCode")
	res := userService.Register(phone, pwd, smsCode, util.SmsRegType)
	c.JSON(200, res)
}

//UpdatePwd 修改密码
func UserUpdatePwd(c *gin.Context) {
	phone := c.PostForm("phone")
	newPwd := c.PostForm("newPwd")
	smsCode := c.PostForm("smsCode")
	res := userService.UpdatePwdService(phone, newPwd, smsCode, util.SmsUpdatePwdType)
	c.JSON(200, res)
}

//ResetPwd 重置密码
func UserResetPwd(c *gin.Context) {
	phone := c.PostForm("phone")
	newPwd := c.PostForm("newPwd")
	smsCode := c.PostForm("smsCode")
	res := userService.ResetPwdService(phone, newPwd, smsCode, util.SmsResetPwdType)
	c.JSON(200, res)
}
