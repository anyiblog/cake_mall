package goods

import (
	"cake_mall/conf"
	"cake_mall/util"
	"time"
)

// GoodsCategory 分类表
type GoodsCategory struct {
	CategoryID   string     `gorm:"primary_key;column:category_id;type:char(36);not null" json:"category_id"` // 分类主键ID
	CategoryName string     `gorm:"column:category_name;type:char(8);not null" json:"category_name"`          // 分类名称
	CreatedAt    time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`                       // 创建时间
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                       // 修改时间
}

//所有分类列表
type ResCategoryList struct {
	Count        string          `json:"count"`
	CategoryList []GoodsCategory `json:"category_list"`
}

func GetCategoryName(CategoryID string) string {
	var categoryName struct {
		CategoryName string
	}
	conf.DB.Model(GoodsCategory{}).Select("category_name").Where("category_id = ?", CategoryID).Scan(&categoryName)
	return categoryName.CategoryName
}

func GetCategoryList(limit, page int) ResCategoryList {
	var resCategoryList ResCategoryList
	conf.DB.Model(GoodsCategory{}).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&resCategoryList.CategoryList)
	conf.DB.Model(GoodsCategory{}).Count(&resCategoryList.Count)
	return resCategoryList
}

func CreateCategory(categoryName string) bool {
	categoryID := util.GenerateUUID()
	c := GoodsCategory{
		CategoryID:   categoryID,
		CategoryName: categoryName,
		CreatedAt:    util.NowTime(),
	}
	if err := conf.DB.Create(c).Error; err != nil {
		return false
	} else {
		return true
	}
}

func UpdateCakeCategory(categoryName, categoryId string) bool {
	data := make(map[string]interface{})
	data["category_name"] = categoryName
	data["updated_at"] = util.NowTime()
	if conf.DB.Model(GoodsCategory{}).Where("category_id = ?", categoryId).Update(data).RecordNotFound() {
		return false
	} else {
		return true
	}
}

func DeleteCakeCategory(categoryId string) bool {
	if conf.DB.Where("category_id = ?", categoryId).Delete(GoodsCategory{}).RecordNotFound() {
		return false
	} else {
		return true
	}
}
