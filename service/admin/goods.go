package admin

import (
	goodsModel "cake_mall/model/goods"
	"cake_mall/serializer"
)

func GetCakeList(limit, page int) serializer.Response {
	basicCakeList := goodsModel.GetBasicSpuList(limit, page)
	return serializer.Response{
		Code: 0,
		Msg:  "获取成功",
		Data: basicCakeList,
	}
}
