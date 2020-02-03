package admin

import (
	adminService "cake_mall/service/admin"
	"cake_mall/util"
	"github.com/gin-gonic/gin"
)

func GetImgListForTag(c *gin.Context) {
	tag := util.StringToInt(c.Query("tag"))
	limit := util.StringToInt(c.Query("limit"))
	page := util.StringToInt(c.Query("page"))
	res := adminService.GetImgListForTag(tag, limit, page)
	c.JSON(200, res)
}

func GetTagList(c *gin.Context) {
	res := adminService.GetTagList()
	c.JSON(200, res)
}
