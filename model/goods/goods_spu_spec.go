package goods

import "time"

// GoodsSpuSpec SPU规格表
type GoodsSpuSpec struct {
	SpuSpecID string    `gorm:"primary_key;column:spu_spec_id;type:char(36);not null" json:"spu_spec_id"` // 产品规格值ID
	SpuID     string    `gorm:"column:spu_id;type:char(36);not null" json:"spu_id"`                       // 产品ID
	SpecID    string    `gorm:"column:spec_id;type:char(36);not null" json:"spec_id"`                     // 规格ID
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}
