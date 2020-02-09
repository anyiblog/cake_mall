package admin

// 后台系统对外API接口

import (
	adminParams "cake_mall/serializer/params/admin"
	adminService "cake_mall/service/admin"
	"github.com/gin-gonic/gin"
)

/**
 * showdoc
 * @catalog 用户
 * @title 登录
 * @description 后台用户登录接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/Login
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param phone 必选 string 手机号
 * @param pwd 必选 string 密码
 * @return {"code": 0,"msg": "登录成功","data": {"token": "BpLnfgDsc3WD9F3qNfHK6a95jjJk+wzDkh0h3+fhfUVuS0jZ9uVbhV4vC6AWX40I","phone": "18325180418","nickname": "安忆","sex": 1,"avatar": "https://wx.qlogo.cn/mmopen/vi_32/AFM9aHX1UicDpzUchB7hIM3Z1v4W0KzZ4IpEWOn1qomxMPKDZEeMTylwAxlbI5QJHUNK2Me3haeDVLK0xtpONpQ/132","balance": 0,"integral": 0,"default_address": ""}}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func UserPwdLogin(c *gin.Context) {
	userLoginParam := &adminParams.UserLoginParam{}
	if err := c.ShouldBind(&userLoginParam); err == nil {
		res := adminService.PwdLogin(userLoginParam.Phone, userLoginParam.Pwd)
		c.JSON(200, res)
	}
}

/**
 * showdoc
 * @catalog 用户
 * @title 退出
 * @description 后台用户退出接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/Logout
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @return {"code":0,"msg":"用户已退出"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func Logout(c *gin.Context) {
	token, _ := c.Get("token")
	res := adminService.Logout(token.(string))
	c.JSON(200, res)
}
