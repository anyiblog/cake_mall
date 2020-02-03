package order

import (
	"time"
)

// OrderInfo 订单信息表
type OrderInfo struct {
	OrderInfoID       string    `gorm:"primary_key;column:order_info_id;type:char(36);not null" json:"order_info_id"`                        // 订单信息ID
	OrderNo           string    `gorm:"column:order_no;type:char(10);not null" json:"order_no"`                                              // 订单编号（10位）
	UserID            string    `gorm:"column:user_id;type:char(36);not null" json:"user_id"`                                                // 用户ID
	OrderCount        int       `gorm:"column:order_count;type:int(11);not null" json:"order_count"`                                         // 订单商品的总数量
	OrderTotalPrice   float64   `gorm:"column:order_total_price;type:decimal(18,2) unsigned zerofill;not null" json:"order_total_price"`     // 订单总额（应付款）
	OrderActualPrice  float64   `gorm:"column:order_actual_price;type:decimal(18,2) unsigned zerofill" json:"order_actual_price"`            // 订单实付款
	OrderFreightPrice float64   `gorm:"column:order_freight_price;type:decimal(18,2) unsigned zerofill;not null" json:"order_freight_price"` // 订单运费
	OrderCreateTime   time.Time `gorm:"column:order_create_time;type:timestamp;not null" json:"order_create_time"`                           // 生成订单时间
	OrderAddressID    string    `gorm:"column:order_address_id;type:char(36);not null" json:"order_address_id"`                              // 订单收货地址（外键ID）
	OrderPayMethod    int       `gorm:"column:order_pay_method;type:int(11)" json:"order_pay_method"`                                        // 订单支付方式
	OrderPayTime      time.Time `gorm:"column:order_pay_time;type:timestamp" json:"order_pay_time"`                                          // 订单支付时间
	OrderDelivery     int       `gorm:"column:order_delivery;type:int(11)" json:"order_delivery"`                                            // 订单配送方式
	OrderStatus       int       `gorm:"column:order_status;type:int(10) unsigned;not null" json:"order_status"`                              // 订单状态
	OrderEndTime      time.Time `gorm:"column:order_end_time;type:timestamp" json:"order_end_time"`                                          // 订单完成时间
	OrderCloseTime    time.Time `gorm:"column:order_close_time;type:timestamp" json:"order_close_time"`                                      // 订单关闭时间
	OtherJSON         string    `gorm:"column:other_json;type:json;not null" json:"other_json"`                                              // 订单备注信息，json格式
	CreatedAt         time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`                                                  // 创建时间
	UpdatedAt         time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                                                  // 更新时间
	DeletedAt         time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                                                  // 删除时间
	DeleteStatus      int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"`                            // 删除状态
}
