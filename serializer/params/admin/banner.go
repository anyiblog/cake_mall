package admin

type CreateBannerParam struct {
	BannerName        string `json:"bannerName" binding:"required"`
	BannerDescription string `json:"bannerDescription" binding:"required"`
	DeletePermiss     int    `json:"deletePermiss" binding:"required"`
}

type CreateBannerItemParam struct {
	ImgUrl         string `json:"imgUrl" binding:"required"`
	BannerId       string `json:"bannerId" binding:"required"`
	BannerItemType int    `json:"bannerItemType" binding:"required"`
}

type DeleteBannerParam struct {
	BannerId string `json:"bannerId" binding:"required"`
}

type UpdateBannerParam struct {
	BannerId          string `json:"bannerId" binding:"required"`
	BannerName        string `json:"bannerName" binding:"required"`
	BannerDescription string `json:"bannerDescription" binding:"required"`
	DeletePermiss     int    `json:"deletePermiss" binding:"required"`
}
