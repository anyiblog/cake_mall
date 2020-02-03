package middleware

import (
	"cake_mall/conf"
	userModel "cake_mall/model/user"
	"cake_mall/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
)

//Token鉴权 每次都检查Redis是否有同等Token，
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("进入token校验")
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(200, serializer.Response{
				Code: serializer.CodeTokenNull,
				Msg:  "请求未携带Token",
			})
			c.Abort()
		} else {
			var UserIdAndToken struct {
				UserId string
				Token  string
			}
			i := conf.DB.Select("user_id,token").Where("token = ?", token).Take(&userModel.User{})
			if i.RowsAffected == 0 {
				c.JSON(200, serializer.Response{
					Code: serializer.CodeTokenIllegalToken,
					Msg:  "非法Token",
				})
				c.Abort()
			} else {
				i.Scan(&UserIdAndToken) //获取数据映射到Struct
				c.Set("token", UserIdAndToken.Token)
				c.Set("userId", UserIdAndToken.UserId)
				c.Next()
			}
		}
	}
}

//Token校验
func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
