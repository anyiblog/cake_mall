package goods

import (
	"cake_mall/conf"
	"time"
)

// GoodsBrand 品牌表
type GoodsBrand struct {
	BrandID      string    `gorm:"primary_key;column:brand_id;type:char(36);not null" json:"brand_id"`       // 品牌ID
	BrandName    string    `gorm:"column:brand_name;type:varchar(50);not null" json:"brand_name"`            // 品牌名称
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`                       // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                       // 更新时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                       // 删除时间
	DeleteStatus int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"` // 删除状态
}

func GetBrandName(BrandID string) string {
	var brandName struct{
		BrandName string
	}
	conf.DB.Model(GoodsBrand{}).Select("brand_name").Where("brand_id = ?",BrandID).Scan(&brandName)
	return brandName.BrandName
}