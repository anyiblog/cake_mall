package goods

import (
	"cake_mall/conf"
	"cake_mall/util"
	"time"
)

// GoodsAttribute 商品属性表
type GoodsAttribute struct {
	AttributeID   string     `gorm:"primary_key;column:attribute_id;type:char(36);not null" json:"attribute_id"` // 属性ID
	AttributeName string     `gorm:"column:attribute_name;type:varchar(50);not null" json:"attribute_name"`      // 属性名称
	CreatedAt     time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`                         // 创建时间
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                         // 更新时间
}

//所有属性列表
type ResCakeAttributeList struct {
	Count             string           `json:"count"`
	CakeAttributeList []GoodsAttribute `json:"cake_attribute_list"`
}

func GetAttributeName(AttributeID string) string {
	var attributeName struct {
		AttributeName string
	}
	conf.DB.Model(GoodsAttribute{}).Select("attribute_name").Where("attribute_id = ?", AttributeID).Scan(&attributeName)
	return attributeName.AttributeName
}

func GetAttributeList(limit, page int) ResCakeAttributeList {
	var resCakeAttributeList ResCakeAttributeList
	conf.DB.Model(GoodsAttribute{}).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&resCakeAttributeList.CakeAttributeList)
	conf.DB.Model(GoodsAttribute{}).Count(&resCakeAttributeList.Count)
	return resCakeAttributeList
}

func CreateCakeAttribute(attributeName string) bool {
	attributeID := util.GenerateUUID()
	at := GoodsAttribute{
		AttributeID:   attributeID,
		AttributeName: attributeName,
		CreatedAt:     util.NowTime(),
	}
	if err := conf.DB.Create(at).Error; err != nil {
		return false
	} else {
		return true
	}
}

func UpdateCakeAttribute(attributeName,attributeId string) bool {
	data := make(map[string]interface{})
	data["attribute_name"] = attributeName
	data["updated_at"] = util.NowTime()
	if conf.DB.Model(GoodsAttribute{}).Where("attribute_id = ?", attributeId).Update(data).RecordNotFound() {
		return false
	} else {
		return true
	}
}

func DeleteCakeAttribute(attributeId string) bool {
	if conf.DB.Where("attribute_id = ?", attributeId).Delete(GoodsAttribute{}).RecordNotFound() {
		return false
	} else {
		return true
	}
}
