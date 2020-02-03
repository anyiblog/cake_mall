package goods

type GoodsImg struct {
	GoodsImgID string `gorm:"primary_key;column:goods_img_id;type:char(36);not null" json:"goods_img_id"`
	ImgID      string `gorm:"column:img_id;type:char(36);not null" json:"img_id"` // 图片表ID
	SpuID      string `gorm:"column:spu_id;type:char(36)" json:"spu_id"`          // 产品ID
	SkuID      string `gorm:"column:sku_id;type:char(36)" json:"sku_id"`          // 商品ID
}