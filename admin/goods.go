package admin

import (
	"cake_mall/serializer"
	adminParams "cake_mall/serializer/params/admin"
	adminService "cake_mall/service/admin"
	"cake_mall/util"
	"github.com/gin-gonic/gin"
)

// 获取甜点列表
/**
 * showdoc
 * @catalog 甜点管理
 * @title 获取甜点列表
 * @description 获取甜点列表接口
 * @method GET
 * @url http://cake.anyiblog.com/admin/CakeList
 * @header token 必选 string 用户token
 * @param limit 必选 string 每页数量
 * @param page 必选 string 页码
 * @return {"code":0,"msg":"获取成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func GetCakeList(c *gin.Context) {
	limit := util.StringToInt(c.Query("limit")) // 每页数量
	page := util.StringToInt(c.Query("page"))   // 页码
	res := adminService.GetCakeList(limit, page)
	c.JSON(200, res)
}

// 获取甜点口味列表
/**
 * showdoc
 * @catalog 甜点管理/甜点口味
 * @title 获取甜点口味列表
 * @description 获取甜点口味列表接口
 * @method GET
 * @url http://cake.anyiblog.com/admin/CakeAttributeList
 * @header token 必选 string 用户token
 * @param limit 必选 string 每页数量
 * @param page 必选 string 页码
 * @return {"code":0,"msg":"获取成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func GetCakeAttributeList(c *gin.Context) {
	limit := util.StringToInt(c.Query("limit")) // 每页数量
	page := util.StringToInt(c.Query("page"))   // 页码
	res := adminService.GetCakeAttributeList(limit, page)
	c.JSON(200, res)
}

//添加甜点口味
/**
 * showdoc
 * @catalog 甜点管理/甜点口味
 * @title 添加甜点口味
 * @description 添加甜点口味接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/CreateCakeAttribute
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param attributeName 必选 string 甜点口味名称
 * @return {"code":0,"msg":"添加成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func CreateCakeAttribute(c *gin.Context) {
	createCakeAttributeParam := &adminParams.CreateCakeAttributeParam{}
	if err := c.ShouldBind(&createCakeAttributeParam); err == nil {
		res := adminService.CreateCakeAttribute(createCakeAttributeParam.AttributeName)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//更新甜点口味
/**
 * showdoc
 * @catalog 甜点管理/甜点口味
 * @title 更新甜点口味
 * @description 更新甜点口味接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/UpdateCakeAttribute
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param attributeName 必选 string 口味名称
 * @param attributeId 必选 string 口味id
 * @return {"code":0,"msg":"更新成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func UpdateCakeAttribute(c *gin.Context) {
	updateCakeAttributeParam := &adminParams.UpdateCakeAttributeParam{}
	if err := c.ShouldBind(&updateCakeAttributeParam); err == nil {
		res := adminService.UpdateCakeAttribute(updateCakeAttributeParam.AttributeName, updateCakeAttributeParam.AttributeId)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//删除甜点口味
/**
 * showdoc
 * @catalog 甜点管理/甜点口味
 * @title 删除甜点口味
 * @description删除甜点口味接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/DeleteCakeAttribute
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param attributeId 必选 string 口味id
 * @return {"code":0,"msg":"删除成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func DeleteCakeAttribute(c *gin.Context) {
	deleteCakeAttributeParam := &adminParams.DeleteCakeAttributeParam{}
	if err := c.ShouldBind(&deleteCakeAttributeParam); err == nil {
		res := adminService.DeleteCakeAttribute(deleteCakeAttributeParam.AttributeId)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// 获取品牌列表
/**
 * showdoc
 * @catalog 甜点管理/品牌
 * @title 获取品牌列表
 * @description 获取品牌列表接口
 * @method GET
 * @url http://cake.anyiblog.com/admin/CakeBrandList
 * @header token 必选 string 用户token
 * @param limit 必选 string 每页数量
 * @param page 必选 string 页码
 * @return {"code":0,"msg":"获取成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func GetCakeBrandList(c *gin.Context) {
	limit := util.StringToInt(c.Query("limit")) // 每页数量
	page := util.StringToInt(c.Query("page"))   // 页码
	res := adminService.GetCakeBrandList(limit, page)
	c.JSON(200, res)
}

// 添加品牌
/**
 * showdoc
 * @catalog 甜点管理/品牌
 * @title 添加品牌
 * @description 添加品牌接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/CreateCakeBrand
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param brandName 必选 string 品牌名称
 * @return {"code":0,"msg":"添加成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func CreateCakeBrand(c *gin.Context) {
	createCakeBrandParam := &adminParams.CreateCakeBrandParam{}
	if err := c.ShouldBind(&createCakeBrandParam); err == nil {
		res := adminService.CreateCakeBrand(createCakeBrandParam.BrandName)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//更新品牌
/**
 * showdoc
 * @catalog 甜点管理/品牌
 * @title 更新品牌
 * @description 更新品牌接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/UpdateCakeBrand
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param brandName 必选 string 品牌名
 * @param brandId 必选 string 品牌id
 * @return {"code":0,"msg":"更新成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func UpdateCakeBrand(c *gin.Context) {
	updateCakeBrandParam := &adminParams.UpdateCakeBrandParam{}
	if err := c.ShouldBind(&updateCakeBrandParam); err == nil {
		res := adminService.UpdateCakeBrand(updateCakeBrandParam.BrandName, updateCakeBrandParam.BrandId)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//删除品牌
/**
 * showdoc
 * @catalog 甜点管理/品牌
 * @title 删除品牌
 * @description 删除品牌接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/DeleteCakeBrand
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param brandId 必选 string 品牌id
 * @return {"code":0,"msg":"删除成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func DeleteCakeBrand(c *gin.Context) {
	deleteCakeBrandParam := &adminParams.DeleteCakeBrandParam{}
	if err := c.ShouldBind(&deleteCakeBrandParam); err == nil {
		res := adminService.DeleteCakeBrand(deleteCakeBrandParam.BrandId)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

// 获取甜点分类列表
/**
 * showdoc
 * @catalog 甜点管理/甜点分类
 * @title 获取甜点分类列表
 * @description 获取甜点分类列表接口
 * @method GET
 * @url http://cake.anyiblog.com/admin/CakeCategoryList
 * @header token 必选 string 用户token
 * @param limit 必选 string 每页数量
 * @param page 必选 string 页码
 * @return {"code":0,"msg":"获取成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func GetCakeCategoryList(c *gin.Context) {
	limit := util.StringToInt(c.Query("limit")) // 每页数量
	page := util.StringToInt(c.Query("page"))   // 页码
	res := adminService.GetCakeCategoryList(limit, page)
	c.JSON(200, res)
}

// 添加甜点分类
/**
 * showdoc
 * @catalog 甜点管理/甜点分类
 * @title 添加甜点分类
 * @description 添加甜点分类接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/CreateCakeCategory
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param categoryName 必选 string 分类名称
 * @return {"code":0,"msg":"添加成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func CreateCakeCategory(c *gin.Context) {
	createCakeCategoryParam := &adminParams.CreateCakeCategoryParam{}
	if err := c.ShouldBind(&createCakeCategoryParam); err == nil {
		res := adminService.CreateCakeCategory(createCakeCategoryParam.CategoryName)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//更新分类
/**
 * showdoc
 * @catalog 甜点管理/甜点分类
 * @title 更新甜点分类
 * @description 更新甜点分类接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/UpdateCakeCategory
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param categoryName 必选 string 分类名称
 * @param categoryId 必选 string 分类id
 * @return {"code":0,"msg":"更新成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func UpdateCakeCategory(c *gin.Context) {
	updateCakeCategoryParam := &adminParams.UpdateCakeCategoryParam{}
	if err := c.ShouldBind(&updateCakeCategoryParam); err == nil {
		res := adminService.UpdateCakeCategory(updateCakeCategoryParam.CategoryName, updateCakeCategoryParam.CategoryId)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}

//删除分类
/**
 * showdoc
 * @catalog 甜点管理/甜点分类
 * @title 删除甜点分类
 * @description 删除甜点分类接口
 * @method POST
 * @url http://cake.anyiblog.com/admin/DeleteCakeCategory
 * @header Content-Type 必选 string application/json
 * @header token 必选 string 用户token
 * @param categoryId 必选 string 分类id
 * @return {"code":0,"msg":"删除成功"}
 * @return_param code int 状态码
 * @return_param msg string 信息
 * @return_param data object 返回数据
 * @remark 无
 */
func DeleteCakeCategory(c *gin.Context) {
	deleteCakeCategoryParam := &adminParams.DeleteCakeCategoryParam{}
	if err := c.ShouldBind(&deleteCakeCategoryParam); err == nil {
		res := adminService.DeleteCakeCategory(deleteCakeCategoryParam.CategoryId)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求参数错误",
			Data: err.Error(),
		})
	}
}
