package model

import (
	"time"
)

// Collect 收藏表
type Collect struct {
	CollectID    string    `gorm:"primary_key;column:collect_id;type:char(36);not null" json:"collect_id"`
	UserID       int       `gorm:"column:user_id;type:int(36) unsigned;not null" json:"user_id"` // 用户ID
	SkuID        int       `gorm:"column:sku_id;type:int(36) unsigned;not null" json:"sku_id"`   // 商品ID
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`           // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                       // 删除时间
	DeleteStatus int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"` // 删除状态
}
