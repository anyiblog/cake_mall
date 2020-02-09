package admin

type UserLoginParam struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
	Pwd   string `form:"pwd" json:"pwd" binding:"required"`
}
