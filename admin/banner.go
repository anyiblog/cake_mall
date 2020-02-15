package admin

import (
	"cake_mall/serializer"
	adminParams "cake_mall/serializer/params/admin"
	adminService "cake_mall/service/admin"
	"github.com/gin-gonic/gin"
)

//创建Banner
/**
 * showdoc
 * @catalog Banner
 * @title 创建Banner
 * @description 创建Banner接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/CreateBanner
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param bannerName 必选 string banner名称
 * @param bannerDescription 必选 string banner描述
 * @param deletePermiss 必选 int 删除权限（0：不可删除，1：可删除）
 * @return {"code":0,"msg":"创建成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func CreateBanner(c *gin.Context) {
	createBannerParam := &adminParams.CreateBannerParam{}
	if err := c.ShouldBind(&createBannerParam); err == nil {
		res := adminService.CreateBanner(createBannerParam.BannerName, createBannerParam.BannerDescription, createBannerParam.DeletePermiss)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//创建Banner子项
/**
 * showdoc
 * @catalog Banner
 * @title 创建Banner子项
 * @description 创建Banner子项接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/CreateBannerItem
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param imgUrl 必选 string 图片url地址
 * @param bannerId 必选 string bannerId
 * @param bannerItemType 必选 int 跳转导向
 * @return {"code":0,"msg":"创建成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func CreateBannerItem(c *gin.Context) {
	createBannerItemParam := &adminParams.CreateBannerItemParam{}
	if err := c.ShouldBind(&createBannerItemParam); err == nil {
		res := adminService.CreateBannerItem(createBannerItemParam.ImgUrl, createBannerItemParam.BannerId, createBannerItemParam.BannerItemType)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// 获取Banner
/**
 * showdoc
 * @catalog Banner
 * @title 获取Banner
 * @description 获取Banner接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/BannerList
 * @header token 必选 string 用户token
 * @return {"code":0,"msg":"获取成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func GetBannerModule(c *gin.Context) {
	res := adminService.GetBannerList()
	c.JSON(200, res)
}

// 删除Banner模块
/**
 * showdoc
 * @catalog Banner
 * @title 删除Banner
 * @description 删除Banner接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/DeleteBanner
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param bannerId 必选 string bannerId
 * @return {"code":0,"msg":"删除成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func DeleteBanner(c *gin.Context) {
	deleteBannerParam := &adminParams.DeleteBannerParam{}
	if err := c.ShouldBind(&deleteBannerParam); err == nil {
		res := adminService.DeleteBanner(deleteBannerParam.BannerId)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//更新Banner
/**
 * showdoc
 * @catalog Banner
 * @title 更新Banner
 * @description 更新Banner接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/UpdateBanner
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param bannerId 必选 string bannerId
 * @param bannerName 必选 string banner名称
 * @param BannerDescription 必选 string banner描述
 * @param DeletePermiss 必选 int 删除权限（0：不可删除，1：可删除）
 * @return {"code":0,"msg":"更新成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func UpdateBanner(c *gin.Context) {
	updateBannerParam := &adminParams.UpdateBannerParam{}
	if err := c.ShouldBind(&updateBannerParam); err == nil {
		res := adminService.UpdateBanner(updateBannerParam.BannerId, updateBannerParam.BannerName, updateBannerParam.BannerDescription, updateBannerParam.DeletePermiss)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}
