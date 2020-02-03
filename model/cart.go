package model

// Cart 购物车表
type Cart struct {
	CartID     string `gorm:"primary_key;column:cart_id;type:char(36);not null" json:"cart_id"`
	UserID     int    `gorm:"column:user_id;type:int(36) unsigned;not null" json:"user_id"` // 用户ID
	SkuID      int    `gorm:"column:sku_id;type:int(36) unsigned;not null" json:"sku_id"`   // 商品ID
	GoodsCount int    `gorm:"column:goods_count;type:int(11);not null" json:"goods_count"`  // 商品数量
}
