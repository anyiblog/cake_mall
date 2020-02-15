package api

// 微信登录or注册
type UserWeChatLoginParam struct {
	Code          string `form:"code" json:"code" binding:"required"`
	EncryptedData string `form:"encryptedData" json:"encryptedData" binding:"required"`
	Iv            string `form:"iv" json:"iv" binding:"required"`
}

// 仅适用第三方登录后无手机信息，绑定手机和密码
type UserBindPhoneAndPwdParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
	Pwd     string `form:"pwd" json:"pwd" binding:"required"`
}

// 验证码登录
type UserSmsLoginParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}

// 密码登录
type UserPwdLoginParam struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
	Pwd   string `form:"pwd" json:"pwd" binding:"required"`
}

// 注册
type UserRegisterParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	Pwd     string `form:"pwd" json:"pwd" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}

// 修改密码
type UserUpdatePwdParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	NewPwd  string `form:"newPwd" json:"newPwd" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}

// 重置密码
type UserResetPwdParam struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	NewPwd  string `form:"newPwd" json:"newPwd" binding:"required"`
	SmsCode string `form:"smsCode" json:"smsCode" binding:"required"`
}
