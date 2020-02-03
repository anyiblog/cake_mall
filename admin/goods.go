package admin

import (
	"cake_mall/util"
	"github.com/gin-gonic/gin"
	adminService "cake_mall/service/admin"
	)

// 获取蛋糕列表
func GetCakeList(c *gin.Context)  {
	limit := util.StringToInt(c.Query("limit")) // 每页数量
	page := util.StringToInt(c.Query("page")) // 页码
	res := adminService.GetCakeList(limit,page)
	c.JSON(200,res)
}