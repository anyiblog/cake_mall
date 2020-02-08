package model

import (
	"cake_mall/conf"
	"cake_mall/util"
	"time"
)

// Banner Banner表
type Banner struct {
	BannerID          string     `gorm:"primary_key;column:banner_id;type:char(36);not null" json:"banner_id"`
	BannerName        string     `gorm:"column:banner_name;type:varchar(50);not null" json:"banner_name"`                // Banner名称，通常作为标识
	BannerDescription string     `gorm:"column:banner_description;type:varchar(255);not null" json:"banner_description"` // Banner描述
	CreatedAt         time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`                             // 创建时间
	UpdatedAt         *time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                             // 更新时间
	DeletePermiss     int        `gorm:"column:delete_permiss;type:int(10) unsigned;not null" json:"delete_permiss"`     // 删除权限(0不可删除，1可删除)
}

type ResBanner struct {
	Count      int      `json:"count"`
	BannerList []Banner `json:"banner_list"`
}

func CreateBanner(BannerName, BannerDescription string, DeletePermiss int) bool {
	bannerID := util.GenerateUUID()
	b := Banner{
		BannerID:          bannerID,
		BannerName:        BannerName,
		BannerDescription: BannerDescription,
		CreatedAt:         util.NowTime(),
		DeletePermiss:     DeletePermiss,
	}
	if err := conf.DB.Create(b).Error; err != nil {
		return false
	} else {
		return true
	}
}

func UpdateBanner(BannerId, BannerName, BannerDescription string, DeletePermiss int) bool {
	data := make(map[string]interface{})
	data["banner_name"] = BannerName
	data["banner_description"] = BannerDescription
	data["delete_permiss"] = DeletePermiss
	if conf.DB.Model(Banner{}).Where("banner_id = ?", BannerId).Update(data).RecordNotFound() {
		return false
	} else {
		return true
	}
}

func GetBanner() ResBanner {
	var banners []Banner
	var count int
	conf.DB.Model(Banner{}).Find(&banners)
	conf.DB.Model(Banner{}).Count(&count)
	return ResBanner{
		Count:      count,
		BannerList: banners,
	}
}

func GetBannerIdForBannerName(BannerName string) string {
	var bannerId struct {
		BannerId string
	}
	conf.DB.Model(Banner{}).Select("banner_id").Where("banner_name = ?", BannerName).Scan(&bannerId)
	return bannerId.BannerId
}

func DeleteBanner(bannerId string) bool {
	if conf.DB.Where("banner_id = ?", bannerId).Delete(Banner{}).RecordNotFound() {
		return false
	} else {
		return true
	}
}
