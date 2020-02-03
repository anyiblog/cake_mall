package order

// OrderGoods 订单商品表
type OrderGoods struct {
	OrderGoodsID string  `gorm:"primary_key;column:order_goods_id;type:char(36);not null" json:"order_goods_id"` // 订单商品ID
	SkuID        string  `gorm:"column:sku_id;type:char(36);not null" json:"sku_id"`                             // 商品ID
	OrderNo      string  `gorm:"column:order_no;type:char(10);not null" json:"order_no"`                         // 订单ID（关联订单信息表）
	UIntPrice    float64 `gorm:"column:uint_price;type:decimal(18,2);not null" json:"uint_price"`                // 订单里商品的单价
	GoodsCount   int     `gorm:"column:goods_count;type:int(11);not null" json:"goods_count"`                    // 订单里商品的数量
	TotalPrice   float64 `gorm:"column:total_price;type:decimal(10,2);not null" json:"total_price"`              // 订单里商品的总价
}