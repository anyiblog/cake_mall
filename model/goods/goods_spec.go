package goods

import "time"

// GoodsSpec 规格表
type GoodsSpec struct {
	SpecID    string    `gorm:"primary_key;column:spec_id;type:char(36);not null" json:"spec_id"` // 规格ID
	SpecName  string    `gorm:"column:spec_name;type:varchar(50);not null" json:"spec_name"`      // 规格名称
	GmtCreate time.Time `gorm:"column:gmt_create;type:timestamp;not null" json:"gmt_create"`
	GmtUpdate time.Time `gorm:"column:gmt_update;type:timestamp;not null" json:"gmt_update"`
}
