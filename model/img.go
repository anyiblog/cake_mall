package model

import (
	"cake_mall/conf"
	"cake_mall/util"
	"time"
)

// Img 图片表
type Img struct {
	ImgID     string    `gorm:"primary_key;column:img_id;type:char(36);not null" json:"img_id"` // 图片主键ID
	ImgURL    string    `gorm:"column:img_url;type:varchar(255);not null" json:"img_url"`       // 图片地址
	ImgTag    int       `gorm:"column:img_tag;type:int(255);not null" json:"img_tag"`           // 图片标签
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`             // 创建时间
}

// 标签信息
type ResTagInfo struct {
	TagId   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

// 图片列表
type ResImgList struct {
	Count   int   `json:"count"`
	ImgList []Img `json:"img_list"`
}

//获取Tag列表
func GetTagList() []ResTagInfo {
	var tag int
	var tagInfo []ResTagInfo
	rows, _ := conf.DB.Debug().Raw("SELECT distinct img_tag from img").Rows()
	for rows.Next() {
		_ = rows.Scan(&tag)
		tagInfo = append(tagInfo, ResTagInfo{
			TagId:   tag,
			TagName: util.ImgDic[tag],
		})
	}
	return tagInfo
}

//根据tag获取图片列表
func GetImgListForTag(tag, limit, page int) ResImgList {
	if page <= 0 {
		return ResImgList{
			Count:   0,
			ImgList: nil,
		}
	} else {
		var resImgList ResImgList
		conf.DB.Debug().Model(Img{}).Where("img_tag = ? ", tag).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&resImgList.ImgList)
		conf.DB.Debug().Model(Img{}).Where("img_tag = ? ", tag).Count(&resImgList.Count)
		return resImgList
	}
}

//图片创建
func CreateImg(imgUrl string) bool {
	imgId := util.GenerateUUID()
	i := Img{
		ImgID:     imgId,
		ImgURL:    imgUrl,
		CreatedAt: util.NowTime(),
	}
	if conf.DB.Create(i).RecordNotFound() {
		return false
	} else {
		return true
	}
}

//图片删除
func DeleteImg(imgUrl string) bool {
	if conf.DB.Where("img_url = ?", imgUrl).Delete(&Img{}).RecordNotFound() {
		return false
	} else {
		return true
	}
}

//从已有网络url的图片创建，不上传COS，返回ImgID
func CreateImgForUrl(imgUrl string) string {
	imgId := util.GenerateUUID()
	i := Img{
		ImgID:     imgId,
		ImgURL:    imgUrl,
		CreatedAt: util.NowTime(),
	}
	conf.DB.Create(i)
	var resImg = struct {
		ImgID  string
		ImgURL string
	}{}
	conf.DB.Model(Img{}).Select("img_id,img_url").Where("img_id = ? ", imgId).Scan(&resImg)
	return resImg.ImgID
}

//获取图片URL
func GetImgUrl(imgId string) string {
	var resImg = struct {
		ImgID  string
		ImgURL string
	}{}
	conf.DB.Model(Img{}).Select("img_id,img_url").Where("img_id = ? ", imgId).Scan(&resImg)
	return resImg.ImgURL
}
