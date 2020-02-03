package order

import "time"

// OrderShipping 物流信息表
type OrderShipping struct {
	ShippingID       string    `gorm:"primary_key;column:shipping_id;type:char(36);not null" json:"shipping_id"`       // 快递信息主键ID
	OrderInfoID      string    `gorm:"column:order_info_id;type:char(36);not null" json:"order_info_id"`               // 订单信息ID（外键）
	ShippingCompName string    `gorm:"column:shipping_comp_name;type:varchar(255);not null" json:"shipping_comp_name"` // 快递公司名称
	ShippingNo       string    `gorm:"column:shipping_no;type:varchar(255);not null" json:"shipping_no"`               // 快递单号
	ShippingTime     time.Time `gorm:"column:shipping_time;type:timestamp;not null" json:"shipping_time"`              // 发货时间
	ReceiveTime      time.Time `gorm:"column:receive_time;type:timestamp;not null" json:"receive_time"`                // 收货时间
	ShippingStatus   int       `gorm:"column:shipping_status;type:int(255);not null" json:"shipping_status"`           // 物流状态
}