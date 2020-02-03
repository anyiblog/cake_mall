package goods

import "time"

// GoodsComment 商品评论表
type GoodsComment struct {
	CommentID      string    `gorm:"primary_key;column:comment_id;type:char(36);not null" json:"comment_id"` // 评论ID
	UserID         string    `gorm:"column:user_id;type:char(36);not null" json:"user_id"`                   // 关联用户ID
	SkuID          string    `gorm:"column:sku_id;type:char(36);not null" json:"sku_id"`                     // 关联商品ID
	SpuID          string    `gorm:"column:spu_id;type:char(36);not null" json:"spu_id"`                     // 关联产品ID
	CommentTime    time.Time `gorm:"column:comment_time;type:timestamp;not null" json:"comment_time"`        // 评论时间
	CommentContent string    `gorm:"column:comment_content;type:longtext;not null" json:"comment_content"`   // 评论内容
	ImgID          string    `gorm:"column:img_id;type:json;not null" json:"img_id"`                         // 图片ID(json)
}
