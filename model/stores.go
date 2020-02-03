package model

// Stores 线下门店表
type Stores struct {
	StoresID          string `gorm:"primary_key;column:stores_id;type:char(36);not null" json:"stores_id"`             // 门店ID
	StoresName        string `gorm:"column:stores_name;type:char(25);not null" json:"stores_name"`                     // 门店名称
	StoresLogo        string `gorm:"column:stores_logo;type:char(36);not null" json:"stores_logo"`                     // 门店logo（关联图片表ID）
	StoresAddress     string `gorm:"column:stores_address;type:varchar(255);not null" json:"stores_address"`           // 门店地址
	StoresContactName string `gorm:"column:stores_contact_name;type:varchar(255);not null" json:"stores_contact_name"` // 门店联系人
	StoresContactTel  string `gorm:"column:stores_contact_tel;type:varchar(255);not null" json:"stores_contact_tel"`   // 门店联系电话
	StoresOpenTime    string `gorm:"column:stores_open_time;type:json;not null" json:"stores_open_time"`               // 门店营业时间
}
