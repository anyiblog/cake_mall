package goods

import "time"

// GoodsSpecValue 规格值表
type GoodsSpecValue struct {
	SpecValueID string    `gorm:"primary_key;column:spec_value_id;type:char(36);not null" json:"spec_value_id"` // 规格值ID
	SpecID      string    `gorm:"column:spec_id;type:char(36);not null" json:"spec_id"`                         // 规格ID
	SpecValue   string    `gorm:"column:spec_value;type:varchar(50);not null" json:"spec_value"`                // 规格值
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}
