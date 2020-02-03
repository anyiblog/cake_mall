package admin

import (
	"cake_mall/model"
	"cake_mall/serializer"
)

//获取Tag列表
func GetTagList() serializer.Response {
	res := model.GetTagList()
	return serializer.Response{
		Code: 0,
		Msg:  "获取成功",
		Data: res,
	}
}

//根据tag获取图片列表
func GetImgListForTag(tag, limit, page int) serializer.Response {
	imgList := model.GetImgListForTag(tag, limit, page)
	return serializer.Response{
		Code: 0,
		Msg:  "获取成功",
		Data: imgList,
	}
}
