package goods

import "time"

// GoodsSkuSpecValue sku规格值
type GoodsSkuSpecValue struct {
	SkuSpecValueID string    `gorm:"primary_key;column:sku_spec_value_id;type:char(36);not null" json:"sku_spec_value_id"` // 商品规格值ID
	SkuID          string    `gorm:"column:sku_id;type:char(36);not null" json:"sku_id"`                                   // 商品ID
	SpecValueID    string    `gorm:"column:spec_value_id;type:char(36);not null" json:"spec_value_id"`                     // 规格值id
	GmtCreate      time.Time `gorm:"column:gmt_create;type:timestamp;not null" json:"gmt_create"`
	GmtUpdate      time.Time `gorm:"column:gmt_update;type:timestamp;not null" json:"gmt_update"`
}
