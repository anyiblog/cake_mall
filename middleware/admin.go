package middleware

import (
	userModel "cake_mall/model/user"
	"cake_mall/serializer"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Get("token")
		if userModel.IsAdminForToken(token.(string)) {
			c.Next()
		} else {
			c.JSON(200, serializer.Response{
				Code: serializer.CodeAuthorizationFailed,
				Msg:  "该用户暂未拥有管理员权限",
				Data: nil,
			})
			c.Abort()
		}
	}
}
