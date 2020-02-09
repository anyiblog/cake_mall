package admin

type ImgTagUpdateParam struct {
	OldTagId int    `json:"oldTagId" binding:"required"`
	NewTagId int    `json:"newTagId" binding:"required"`
	ImgId    string `json:"imgId" binding:"required"`
}
