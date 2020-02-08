package admin

import (
	adminService "cake_mall/service/admin"
	"cake_mall/util"
	"github.com/gin-gonic/gin"
)

//创建Banner
func CreateBanner(c *gin.Context) {
	bannerName := c.PostForm("bannerName")
	bannerDescription := c.PostForm("bannerDescription")
	deletePermiss := util.StringToInt(c.PostForm("deletePermiss"))
	res := adminService.CreateBanner(bannerName, bannerDescription, deletePermiss)
	c.JSON(200, res)
}

//创建Banner子项
func CreateBannerItem(c *gin.Context) {
	imgUrl := c.PostForm("imgUrl")
	bannerId := c.PostForm("bannerId")
	bannerItemType := util.StringToInt(c.PostForm("bannerItemType"))
	res := adminService.CreateBannerItem(imgUrl, bannerId, bannerItemType)
	c.JSON(200, res)
}

// 获取Banner模块
func GetBannerModule(c *gin.Context) {
	res := adminService.GetBannerList()
	c.JSON(200, res)
}

// 删除Banner模块
func DeleteBanner(c *gin.Context) {
	bannerId := c.PostForm("bannerId")
	res := adminService.DeleteBanner(bannerId)
	c.JSON(200, res)
}

//更新Banner模块
func UpdateBanner(c *gin.Context) {
	bannerId := c.PostForm("bannerId")
	bannerName := c.PostForm("bannerName")
	bannerDescription := c.PostForm("bannerDescription")
	deletePermiss := util.StringToInt(c.PostForm("deletePermiss"))
	res := adminService.UpdateBanner(bannerId, bannerName, bannerDescription, deletePermiss)
	c.JSON(200, res)
}
