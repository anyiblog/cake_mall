package goods

import (
	"cake_mall/conf"
	"time"
)

// GoodsCategory 分类表
type GoodsCategory struct {
	CategoryID   string    `gorm:"primary_key;column:category_id;type:char(36);not null" json:"category_id"` // 分类主键ID
	CategoryName string    `gorm:"column:category_name;type:char(8);not null" json:"category_name"`          // 分类名称
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`                       // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                       // 修改时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                       // 删除时间
	DeleteStatus int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"` // 删除状态
}

func GetCategoryName(CategoryID string) string {
	var categoryName struct{
		CategoryName string
	}
	conf.DB.Model(GoodsCategory{}).Select("category_name").Where("category_id = ?",CategoryID).Scan(&categoryName)
	return categoryName.CategoryName
}
