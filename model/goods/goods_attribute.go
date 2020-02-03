package goods

import (
	"cake_mall/conf"
	"time"
)

// GoodsAttribute 商品属性表
type GoodsAttribute struct {
	AttributeID   string    `gorm:"primary_key;column:attribute_id;type:char(36);not null" json:"attribute_id"` // 属性ID
	AttributeName string    `gorm:"column:attribute_name;type:varchar(50);not null" json:"attribute_name"`      // 属性名称
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`                         // 创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                         // 更新时间
	DeletedAt     time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                         // 删除时间
	DeleteStatus  int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"`   // 删除状态
}

func GetAttributeName(AttributeID string) string {
	var attributeName struct{
		AttributeName string
	}
	conf.DB.Model(GoodsAttribute{}).Select("attribute_name").Where("attribute_id = ?",AttributeID).Scan(&attributeName)
	return attributeName.AttributeName
}
