package goods

import (
	"time"
)

// GoodsSku Sku表（具体的某个商品）
type GoodsSku struct {
	SkuID            string    `gorm:"primary_key;column:sku_id;type:char(36);not null" json:"sku_id"`           // 商品主键ID
	SkuNo            string    `gorm:"column:sku_no;type:char(11);not null" json:"sku_no"`                       // 商品编号，唯一（11位）
	SpuID            string    `gorm:"column:spu_id;type:char(36);not null" json:"spu_id"`                       // 产品ID（spuid）
	SkuName          string    `gorm:"column:sku_name;type:char(10);not null" json:"sku_name"`                   // 商品名称（用于连接SPU名称）
	SkuImg           string    `gorm:"column:sku_img;type:char(36);not null" json:"sku_img"`                     // 商品图片（imgId）
	SkuPrice         float64   `gorm:"column:sku_price;type:decimal(18,2);not null" json:"sku_price"`            // 商品售价
	SkuLowPrice      float64   `gorm:"column:sku_low_price;type:decimal(18,2);not null" json:"sku_low_price"`    // 商品最低售价
	SkuSales         int       `gorm:"column:sku_sales;type:int(11) unsigned;not null" json:"sku_sales"`         // 商品销量
	SkuStock         int       `gorm:"column:sku_stock;type:int(11);not null" json:"sku_stock"`                  // 商品库存
	SkuPostage       float64   `gorm:"column:sku_postage;type:decimal(18,2);not null" json:"sku_postage"`        // 配送费（0为免费）
	SkuAttributeJSON string    `gorm:"column:sku_attribute_json;type:json;not null" json:"sku_attribute_json"`   // 其他信息
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`                       // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                       // 更新时间
	DeletedAt        time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                       // 删除时间
	DeleteStatus     int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"` // 删除状态
}