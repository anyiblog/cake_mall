package goods

// GoodsDetails 商品详情表
type GoodsDetails struct {
	DetailsID      string `gorm:"primary_key;column:details_id;type:char(36);not null" json:"details_id"` // 详情ID
	SpuID          string `gorm:"column:spu_id;type:char(36);not null" json:"spu_id"`                     // 对应的产品ID
	DetailsContent string `gorm:"column:details_content;type:longtext;not null" json:"details_content"`   // 内容
}
