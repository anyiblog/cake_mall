package util

const (
	//第三方平台登录

	OpenTypeWeChat = 1 //1 为微信登录

	//短信验证码Type

	SmsRegType       = 527632 //注册
	SmsLoginType     = 527633 //登录
	SmsBindPhoneType = 527634 //绑定手机号
	SmsResetPwdType  = 527635 //重置密码
	SmsUpdatePwdType = 527636 // 修改密码

	//Redis库

	RedisToken      = 0
	RedisUserInfo   = 1
	RedisWeChatUser = 2
	RedisSms        = 3
)

var ImgDic = map[int]string{
	0: "所有图片",
	1: "头像图片",
	2: "商品图片",
	3: "其他图片",
}
