package order

// OrderSelfPick 订单自提表
type OrderSelfPick struct {
	SelfPickID     string `gorm:"primary_key;column:self_pick_id;type:char(36);not null" json:"self_pick_id"` // 门店自取表主键ID
	OrderInfoID    string `gorm:"column:order_info_id;type:char(36);not null" json:"order_info_id"`           // 订单信息关联ID
	UserID         string `gorm:"column:user_id;type:char(36);not null" json:"user_id"`                       // 用户信息关联ID
	StoresID       string `gorm:"column:stores_id;type:char(36);not null" json:"stores_id"`                   // 自提门店关联ID
	SelfPickCode   string `gorm:"column:self_pick_code;type:varchar(12);not null" json:"self_pick_code"`      // 自提码
	SelfPickName   string `gorm:"column:self_pick_name;type:varchar(255);not null" json:"self_pick_name"`     // 提货人姓名
	SelfPickTel    string `gorm:"column:self_pick_tel;type:varchar(11);not null" json:"self_pick_tel"`        // 提货人手机号码，短信验证码使用
	SelfPickStatus int    `gorm:"column:self_pick_status;type:int(11);not null" json:"self_pick_status"`      // 自取状态
}
