package admin

import (
	adminParams "cake_mall/serializer/params/admin"
	adminService "cake_mall/service/admin"
	"cake_mall/util"
	"github.com/gin-gonic/gin"
)

//根据tagId获取图片列表
/**
 * showdoc
 * @catalog 媒体库
 * @title 获取图片列表
 * @description 获取图片列表接口
 * @method GET
 * @url http://cake.anyiblog.com/admin/ImgList
 * @header token 必选 string 用户token
 * @param tag 必选 int tagId
 * @param limit 必选 int 每页数量
 * @param page 必选 string 页码
 * @return {"code":0,"msg":"获取成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func GetImgListForTag(c *gin.Context) {
	tag := util.StringToInt(c.Query("tag"))
	limit := util.StringToInt(c.Query("limit"))
	page := util.StringToInt(c.Query("page"))
	res := adminService.GetImgListForTag(tag, limit, page)
	c.JSON(200, res)
}

//获取tag列表
/**
 * showdoc
 * @catalog 媒体库
 * @title 获取tag列表
 * @description 获取tag列表接口
 * @method GET
 * @url http://cake.anyiblog.com/admin/TagList
 * @header token 必选 string 用户token
 * @return {"code":0,"msg":"获取成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func GetTagList(c *gin.Context) {
	res := adminService.GetTagList()
	c.JSON(200, res)
}

// 更新图片tag
/**
 * showdoc
 * @catalog 媒体库
 * @title 更新图片tag
 * @description 更新图片tag接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/ImgTagUpdate
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param oldTagId 必选 int 旧tagId
 * @param newTagId 必选 int 新tagId
 * @param imgId 必选 string 图片id
 * @return {"code":0,"msg":"更新成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func ImgTagUpdate(c *gin.Context) {
	imgTagUpdateParam := &adminParams.ImgTagUpdateParam{}
	if err := c.ShouldBind(&imgTagUpdateParam); err == nil {
		res := adminService.ImgTagUpdate(imgTagUpdateParam.OldTagId, imgTagUpdateParam.NewTagId, imgTagUpdateParam.ImgId)
		c.JSON(200, res)
	}
}
