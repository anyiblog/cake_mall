package model

import "time"

// Banner Banner表
type Banner struct {
	BannerID          string    `gorm:"primary_key;column:banner_id;type:char(36);not null" json:"banner_id"`
	BannerName        string    `gorm:"column:banner_name;type:varchar(50);not null" json:"banner_name"`                // Banner名称，通常作为标识
	BannerDescription string    `gorm:"column:banner_description;type:varchar(255);not null" json:"banner_description"` // Banner描述
	DeletedAt         time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                             // 删除时间
	UpdatedAt         time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                             // 更新时间
	DeleteStatus      int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"`       // 删除状态
	DeletePermiss     int       `gorm:"column:delete_permiss;type:int(10) unsigned;not null" json:"delete_permiss"`     // 删除权限(0不可删除，1可删除)
}