package admin

import (
	"cake_mall/model"
	"cake_mall/serializer"
)

//创建
func CreateBanner(BannerName, BannerDescription string, DeletePermiss int) serializer.Response {
	if model.CreateBanner(BannerName, BannerDescription, DeletePermiss) {
		return serializer.Response{
			Code: 0,
			Msg:  "创建Banner成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "创建Banner失败",
		}
	}
}
func DeleteBanner(bannerId string) serializer.Response {
	if model.DeleteBanner(bannerId) {
		return serializer.Response{
			Code: 0,
			Msg:  "删除Banner成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "删除Banner失败",
		}
	}
}
func UpdateBanner(BannerId, BannerName, BannerDescription string, DeletePermiss int) serializer.Response {
	if model.UpdateBanner(BannerId, BannerName, BannerDescription, DeletePermiss) {
		return serializer.Response{
			Code: 0,
			Msg:  "更新BannerItem成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "更新BannerItem失败",
		}
	}

}

func CreateBannerItem(ImgUrl, BannerId string, BannerItemType int) serializer.Response {
	if model.CreateBannerItem(ImgUrl, BannerId, BannerItemType) {
		return serializer.Response{
			Code: 0,
			Msg:  "创建BannerItem成功",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "创建BannerItem失败",
		}
	}
}

func GetBannerList() serializer.Response {
	return serializer.Response{
		Code: 0,
		Msg:  "获取成功",
		Data: model.GetBanner(),
	}
}
