package goods

import (
	"cake_mall/conf"
	"cake_mall/util"
	"time"
)

// GoodsBrand 品牌表
type GoodsBrand struct {
	BrandID   string     `gorm:"primary_key;column:brand_id;type:char(36);not null" json:"brand_id"` // 品牌ID
	BrandName string     `gorm:"column:brand_name;type:varchar(50);not null" json:"brand_name"`      // 品牌名称
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`                 // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                 // 更新时间
}

//所有品牌列表
type ResBrandList struct {
	Count     string       `json:"count"`
	BrandList []GoodsBrand `json:"brand_list"`
}

func GetBrandName(BrandID string) string {
	var brandName struct {
		BrandName string
	}
	conf.DB.Model(GoodsBrand{}).Select("brand_name").Where("brand_id = ?", BrandID).Scan(&brandName)
	return brandName.BrandName
}

func GetCakeBrandList(limit, page int) ResBrandList {
	var resBrandList ResBrandList
	conf.DB.Model(GoodsAttribute{}).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&resBrandList.BrandList)
	conf.DB.Model(GoodsAttribute{}).Count(&resBrandList.Count)
	return resBrandList
}

func CreateCakeBrand(brandName string) bool {
	branID := util.GenerateUUID()
	b := GoodsBrand{
		BrandID:   branID,
		BrandName: brandName,
		CreatedAt: util.NowTime(),
	}
	if err := conf.DB.Create(b).Error; err != nil {
		return false
	} else {
		return true
	}
}

func UpdateCakeBrand(brandName,brandId string) bool {
	data := make(map[string]interface{})
	data["brand_name"] = brandName
	data["updated_at"] = util.NowTime()
	if conf.DB.Model(GoodsBrand{}).Where("brand_id = ?", brandId).Update(data).RecordNotFound() {
		return false
	} else {
		return true
	}
}

func DeleteCakeBrand(brandId string) bool {
	if conf.DB.Where("brand_id = ?", brandId).Delete(GoodsBrand{}).RecordNotFound() {
		return false
	} else {
		return true
	}
}
