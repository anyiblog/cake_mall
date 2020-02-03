package util

const (
	//第三方平台登录

	OpenTypeWeChat = 1 //1 为微信登录

	//短信验证码Type

	SmsRegType       = 484081 //注册
	SmsLoginType     = 484085 //登录
	SmsBindPhoneType = 506143 //绑定手机号
	SmsResetPwdType  = 508003 //重置密码
	SmsUpdatePwdType = 508008 // 修改密码

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
