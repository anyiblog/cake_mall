package model

// BannerItem Banner子项表
type BannerItem struct {
	BannerItemID   string `gorm:"primary_key;column:banner_item_id;type:char(36);not null" json:"banner_item_id"`
	ImgID          string `gorm:"column:img_id;type:char(36)" json:"img_id"`                             // 外键图片ID
	BannerID       string `gorm:"column:banner_id;type:char(36)" json:"banner_id"`                       // 外键BannerID
	BannerItemType int    `gorm:"column:banner_item_type;type:int(11);not null" json:"banner_item_type"` // 跳转类型，0 无导向，1 商品， 2 专题， 3 外部链接
}
