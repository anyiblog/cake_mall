package goods

import (
	"cake_mall/conf"
	"cake_mall/util"
	"fmt"
	"time"
)

// GoodsSpu Spu表（产品表）
type GoodsSpu struct {
	SpuID              string    `gorm:"primary_key;column:spu_id;type:char(36);not null" json:"spu_id"`                   // 产品主键ID
	SpuNo              string    `gorm:"column:spu_no;type:char(11);not null" json:"spu_no"`                               // 产品编号,唯一(11位)
	SpuName            string    `gorm:"column:spu_name;type:varchar(50);not null" json:"spu_name"`                        // 产品名称
	SpuRecommendReason string    `gorm:"column:spu_recommend_reason;type:varchar(255)" json:"spu_recommend_reason"`        // 产品推荐标语
	SpuImg             string    `gorm:"column:spu_img;type:char(36);not null" json:"spu_img"`                             // 产品图片（imgID）唯一
	SpuPrice           float64   `gorm:"column:spu_price;type:decimal(18,2);not null" json:"spu_price"`                    // 产品售价
	SpuShippingMethod  string    `gorm:"column:spu_shipping_method;type:varchar(255);not null" json:"spu_shipping_method"` // 产品配送方式
	SpuIbs             string    `gorm:"column:spu_ibs;type:json" json:"spu_ibs"`                                          // 产品原材料
	SpuAttributeJSON   string    `gorm:"column:spu_attribute_json;type:json;not null" json:"spu_attribute_json"`           // 产品备注信息
	ImgBanner          string    `gorm:"column:img_banner;type:json;not null" json:"img_banner"`                           // banner图片列表（json）
	CategoryID         string    `gorm:"column:category_id;type:char(36)" json:"category_id"`                              // 分类ID
	BrandID            string    `gorm:"column:brand_id;type:char(36);not null" json:"brand_id"`                           // 品牌ID
	AttributeID        string    `gorm:"column:attribute_id;type:char(36)" json:"attribute_id"`                            // 属性ID
	CreatedAt          time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`                               // 创建时间
	UpdatedAt          time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                               // 更新时间
	DeletedAt          time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                               // 删除时间
	DeleteStatus       int       `gorm:"column:delete_status;type:int(10) unsigned;not null" json:"delete_status"`         // 删除状态
}

type ResBasicCakeInfo struct {
	Count    int             `json:"count"`
	CakeList []BasicCakeInfo `json:"cake_list"`
}

type BasicCakeInfo struct {
	ID             string    `json:"id"`
	No             string    `json:"no"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
	ShippingMethod string    `json:"shippingMethod"`
	Category       string    `json:"category"`
	Brand          string    `json:"brand"`
	Attribute      string    `json:"attribute"`
	Status         int       `json:"status"`
	ReleaseTime    time.Time `json:"releaseTime"`
}

// 获取基础信息列表
func GetBasicSpuList(limit, page int) ResBasicCakeInfo {
	var cakeList []BasicCakeInfo
	allSpuInfo, count := GetAllSpuInfo(limit, page)
	for _, value := range allSpuInfo {
		categoryName := GetCategoryName(value.CategoryID)
		brandName := GetBrandName(value.BrandID)
		attributeName := GetAttributeName(value.AttributeID)
		cakeList = append(cakeList, BasicCakeInfo{
			ID:             value.SpuID,
			No:             value.SpuNo,
			Name:           value.SpuName,
			Price:          value.SpuPrice,
			ShippingMethod: value.SpuShippingMethod,
			Category:       categoryName,
			Brand:          brandName,
			Attribute:      attributeName,
			Status:         value.DeleteStatus,
			ReleaseTime:    util.TimeFormat(value.CreatedAt),
		})
	}
	fmt.Println(cakeList)
	resBasicCakeInfo := ResBasicCakeInfo{
		Count:    count,
		CakeList: cakeList,
	}
	return resBasicCakeInfo
}

// 获取所有SPU信息，返回一个数组
func GetAllSpuInfo(limit, page int) (SpuInfo []GoodsSpu, count int) {
	var allSpuInfo []GoodsSpu
	conf.DB.Model(GoodsSpu{}).Debug().Order("spu_id desc").Limit(limit).Offset((page - 1) * limit).Find(&allSpuInfo)
	conf.DB.Model(GoodsSpu{}).Count(&count)
	return allSpuInfo, count
}
