package admin

// 后台系统对外API接口

import (
	"cake_mall/serializer"
	adminService "cake_mall/service/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

//UserPwdLogin 密码登录
func UserPwdLogin(c *gin.Context) {
	//phone := c.PostForm("phone")
	//pwd := c.PostForm("pwd")
	var paramUserLogin struct {
		Phone string `json:"phone" binding:"required"`
		Pwd   string `json:"pwd" binding:"required"`
	}

	if err := c.ShouldBind(&paramUserLogin); err == nil {
		res := adminService.PwdLogin(paramUserLogin.Phone, paramUserLogin.Pwd)
		c.JSON(200, res)
	}
}

func UserInfo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "用户信息",
		Data: nil,
	})
}

//退出登录
func Logout(c *gin.Context) {
	token, _ := c.Get("token")
	fmt.Println(token)
	res := adminService.Logout(token.(string))
	c.JSON(200, res)
}
