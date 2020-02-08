package model

import (
	"cake_mall/conf"
	"cake_mall/util"
)

// BannerItem Banner子项表
type BannerItem struct {
	BannerItemID   string `gorm:"primary_key;column:banner_item_id;type:char(36);not null" json:"banner_item_id"`
	ImgUrl         string `gorm:"column:img_url;type:varchar(255)" json:"img_url"`                       // 图片链接
	BannerID       string `gorm:"column:banner_id;type:char(36)" json:"banner_id"`                       // 外键BannerID
	BannerItemType int    `gorm:"column:banner_item_type;type:int(11);not null" json:"banner_item_type"` // 跳转类型，0 无导向，1 商品， 2 专题， 3 外部链接
}

func CreateBannerItem(ImgUrl, BannerId string, BannerItemType int) bool {
	bannerItemID := util.GenerateUUID()
	bi := BannerItem{
		BannerItemID:   bannerItemID,
		ImgUrl:         ImgUrl,
		BannerID:       BannerId,
		BannerItemType: BannerItemType,
	}
	if err := conf.DB.Create(bi).Error; err != nil {
		return false
	} else {
		return true
	}
}
