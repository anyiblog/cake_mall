package admin

import (
	goodsModel "cake_mall/model/goods"
	"cake_mall/serializer"
)

// 获取甜点列表
func GetCakeList(limit, page int) serializer.Response {
	basicCakeList := goodsModel.GetBasicSpuList(limit, page)
	return serializer.Response{
		Code: 0,
		Msg:  "获取甜点列表成功",
		Data: basicCakeList,
	}
}

// 获取甜点口味列表
func GetCakeAttributeList(limit, page int) serializer.Response {
	cakeAttributeList := goodsModel.GetAttributeList(limit, page)
	return serializer.Response{
		Code: 0,
		Msg:  "获取甜点口味列表成功",
		Data: cakeAttributeList,
	}
}

// 添加甜点口味
func CreateCakeAttribute(attributeName string) serializer.Response {
	if goodsModel.CreateCakeAttribute(attributeName) {
		return serializer.Response{
			Code: 0,
			Msg:  "添加甜点口味成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "添加甜点口味失败",
		}
	}
}

// 修改甜点口味
func UpdateCakeAttribute(attributeName, attributeId string) serializer.Response {
	if goodsModel.UpdateCakeAttribute(attributeName, attributeId) {
		return serializer.Response{
			Code: 0,
			Msg:  "修改甜点口味信息成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "修改甜点口味信息失败",
		}
	}
}

// 删除甜点口味
func DeleteCakeAttribute(attributeId string) serializer.Response {
	if goodsModel.DeleteCakeAttribute(attributeId) {
		return serializer.Response{
			Code: 0,
			Msg:  "删除甜点口味成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "删除甜点口味失败",
		}
	}
}

// 获取品牌列表
func GetCakeBrandList(limit, page int) serializer.Response {
	brandList := goodsModel.GetCakeBrandList(limit, page)
	return serializer.Response{
		Code: 0,
		Msg:  "获取品牌列表成功",
		Data: brandList,
	}
}

// 添加品牌
func CreateCakeBrand(brandName string) serializer.Response {
	if goodsModel.CreateCakeBrand(brandName) {
		return serializer.Response{
			Code: 0,
			Msg:  "添加品牌成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "添加品牌失败",
		}
	}
}

// 修改品牌
func UpdateCakeBrand(brandName, brandId string) serializer.Response {
	if goodsModel.UpdateCakeBrand(brandName, brandId) {
		return serializer.Response{
			Code: 0,
			Msg:  "修改品牌信息成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "修改品牌信息失败",
		}
	}
}

// 删除品牌
func DeleteCakeBrand(brandId string) serializer.Response {
	if goodsModel.DeleteCakeBrand(brandId) {
		return serializer.Response{
			Code: 0,
			Msg:  "删除品牌成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "删除品牌失败",
		}
	}
}

// 获取甜点分类列表
func GetCakeCategoryList(limit, page int) serializer.Response {
	brandList := goodsModel.GetCategoryList(limit, page)
	return serializer.Response{
		Code: 0,
		Msg:  "获取甜点分类列表成功",
		Data: brandList,
	}
}

// 添加甜点分类
func CreateCakeCategory(categoryName string) serializer.Response {
	if goodsModel.CreateCategory(categoryName) {
		return serializer.Response{
			Code: 0,
			Msg:  "添加分类成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "添加分类失败",
		}
	}
}

// 修改甜点分类
func UpdateCakeCategory(categoryName, categoryId string) serializer.Response {
	if goodsModel.UpdateCakeCategory(categoryName, categoryId) {
		return serializer.Response{
			Code: 0,
			Msg:  "修改甜点分类信息成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "修改甜点分类信息失败",
		}
	}
}

// 删除甜点分类
func DeleteCakeCategory(categoryId string) serializer.Response {
	if goodsModel.DeleteCakeCategory(categoryId) {
		return serializer.Response{
			Code: 0,
			Msg:  "删除甜点分类成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "删除甜点分类失败",
		}
	}
}