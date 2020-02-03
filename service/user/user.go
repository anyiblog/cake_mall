package user

import (
	userModel "cake_mall/model/user"
	"cake_mall/serializer"
	"cake_mall/service"
	"cake_mall/util"
	"cake_mall/util/httplib"
	"fmt"
	"net/url"
	"os"
)

//微信登录逻辑
func WeChatLoginService(jsCode string, encryptedData string, iv string) serializer.Response {
	host, _ := url.Parse("https://api.weixin.qq.com/sns/jscode2session?")
	valueList, _ := url.ParseQuery("")
	valueList.Add("appid", os.Getenv("WeChat_Appid"))
	valueList.Add("secret", os.Getenv("WeChat_Secret"))
	valueList.Add("js_code", jsCode)
	valueList.Add("grant_type", "authorize")
	queryStr := valueList.Encode()
	reqUrl := host.String() + queryStr //构造后的URL（string）
	res := httplib.Get(reqUrl)
	resStrJson, _ := res.String()
	weChatRes := util.StrJsonToMap(resStrJson) //string类型的json转map
	sessionKey := weChatRes["session_key"]
	wxUserInfo := util.WxDataDecrypt(sessionKey, encryptedData, iv)
	fmt.Println(wxUserInfo.OpenID)
	if weChatRes["openid"] == wxUserInfo.OpenID {
		userId := userModel.CreateWxUser(wxUserInfo)
		resUserInfo := userModel.GetUserInfo(userId)
		service.SaveTokenRedis(userId, resUserInfo.Token)
		return serializer.Response{
			Msg:  "登录成功",
			Data: resUserInfo,
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "登录失败",
		}
	}
}

//绑定手机和密码
func BindPhoneAndPwdService(userId, phone, smsCode, pwd string, smsType int) serializer.Response {
	if service.CheckSmsRedis(phone, util.IntToString(smsType), smsCode) { //检验验证是否一致
		if userModel.BindPhoneAndPwd(userId, phone, pwd) {
			return serializer.Response{
				Msg: "绑定成功",
			}
		} else {
			return serializer.Response{
				Code: 1,
				Msg:  "绑定失败，请联系客服",
			}
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "验证码不一致",
		}
	}
}

//验证码登录
func SmsLogin(phone, smsCode string, smsType int) serializer.Response {
	if userModel.IsExist(phone) { //校验是否存在手机号相关用户
		if service.CheckSmsRedis(phone, util.IntToString(smsType), smsCode) { //检验验证是否一致
			userId := userModel.GetUserIdForPhone(phone)
			if userId != "" {
				resUserInfo := userModel.GetUserInfo(userId)
				return serializer.Response{
					Msg:  "登录成功",
					Data: resUserInfo,
				}
			} else {
				return serializer.Response{
					Code: 1,
					Msg:  "用户名密码错误",
				}
			}
		} else {
			return serializer.Response{
				Code: 1,
				Msg:  "验证码不一致",
			}
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "账号未注册",
		}
	}
}

//密码登录
func PwdLogin(phone, pwd string) serializer.Response {
	if userModel.IsExist(phone) {
		userId := userModel.GetUserIdForPwd(phone, pwd)
		if userId != "" {
			resUserInfo := userModel.GetUserInfo(userId)
			return serializer.Response{
				Msg:  "登录成功",
				Data: resUserInfo,
			}
		} else {
			return serializer.Response{
				Code: 1,
				Msg:  "用户名密码错误",
			}
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "账号未注册",
		}
	}
}

func Register(phone, pwd, smsCode string, smsType int) serializer.Response {
	if !userModel.IsExist(phone) { //校验是否存在手机号相关用户
		if service.CheckSmsRedis(phone, util.IntToString(smsType), smsCode) { //检验验证是否一致
			userId := userModel.CreateUser(phone, pwd)
			if userId != "" {
				resUserInfo := userModel.GetUserInfo(userId)
				return serializer.Response{
					Msg:  "注册成功",
					Data: resUserInfo,
				}
			} else {
				return serializer.Response{
					Code: 1,
					Msg:  "注册失败，请联系客服",
				}
			}
		} else {
			return serializer.Response{
				Code: 1,
				Msg:  "验证码不一致",
			}
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "账号已存在",
		}
	}
}

//更新密码
func UpdatePwdService(phone, newPwd, smsCode string, smsType int) serializer.Response {
	if service.CheckSmsRedis(phone, util.IntToString(smsType), smsCode) { //检验验证是否一致
		if userModel.UpdatePwd(phone, newPwd) { //更新密码
			return serializer.Response{
				Msg: "密码更新成功",
			}
		} else {
			return serializer.Response{
				Code: 1,
				Msg:  "密码更新失败，请联系客服",
			}
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "验证码不一致",
		}
	}
}

//重置密码
func ResetPwdService(phone, newPwd, smsCode string, smsType int) serializer.Response {
	if userModel.IsExist(phone) { //校验是否存在手机号相关用户
		if service.CheckSmsRedis(phone, util.IntToString(smsType), smsCode) { //检验验证是否一致
			if userModel.UpdatePwd(phone, newPwd) { //更新密码
				return serializer.Response{
					Msg: "密码重置成功",
				}
			} else {
				return serializer.Response{
					Code: 1,
					Msg:  "密码重置失败，请联系客服",
				}
			}
		} else {
			return serializer.Response{
				Code: 1,
				Msg:  "验证码不一致",
			}
		}
	}else{
		return serializer.Response{
			Code: 1,
			Msg:  "手机未注册",
		}
	}
}
