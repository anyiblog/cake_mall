package user

import "cake_mall/conf"

// UserAddress 用户收货地址表
type UserAddress struct {
	AddressID       string `gorm:"primary_key;column:address_id;type:char(36);not null" json:"address_id"`     // 地址ID
	UserID          string `gorm:"column:user_id;type:char(36);not null" json:"user_id"`                       // 用户ID
	Consignee       string `gorm:"column:consignee;type:varchar(20);not null" json:"consignee"`                // 收货人
	Tel             string `gorm:"column:tel;type:char(11);not null" json:"tel"`                               // 电话
	Province        string `gorm:"column:province;type:varchar(20);not null" json:"province"`                  // 省
	City            string `gorm:"column:city;type:varchar(20);not null" json:"city"`                          // 市
	County          string `gorm:"column:county;type:varchar(20);not null" json:"county"`                      // 县
	Street          string `gorm:"column:street;type:varchar(120);not null" json:"street"`                     // 街道地址
	DetailedAddress string `gorm:"column:detailed_address;type:varchar(255);not null" json:"detailed_address"` // 详细地址
}

func CreateAddress(userId , addressInfo string)  {

}

func GetDetailAddress(userId string) string {
	var resAddress struct{
		DetailedAddress string
	}
	conf.DB.Model(UserAddress{}).Select("detailed_address").Where("user_id = ?",userId).Scan(&resAddress)
	return resAddress.DetailedAddress
}
