package admin

import (
	adminService "cake_mall/service/admin"
	"cake_mall/util"
	"github.com/gin-gonic/gin"
)

//获取Tag列表
func GetImgListForTag(c *gin.Context) {
	tag := util.StringToInt(c.Query("tag"))
	limit := util.StringToInt(c.Query("limit"))
	page := util.StringToInt(c.Query("page"))
	res := adminService.GetImgListForTag(tag, limit, page)
	c.JSON(200, res)
}

//根据tag获取图片列表
func GetTagList(c *gin.Context) {
	res := adminService.GetTagList()
	c.JSON(200, res)
}

//更新图片tag
func ImgTagUpdate(c *gin.Context) {
	oldTagId := util.StringToInt(c.PostForm("oldTagId"))
	newTagId := util.StringToInt(c.PostForm("newTagId"))
	imgId := c.PostForm("imgId")
	res := adminService.ImgTagUpdate(oldTagId, newTagId, imgId)
	c.JSON(200, res)
}
