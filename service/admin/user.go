package admin

// 后台管理系统逻辑层

import (
	userModel "cake_mall/model/user"
	"cake_mall/serializer"
	"cake_mall/service"
)

// 用户登录
func PwdLogin(phone, pwd string) serializer.Response {
	if userModel.IsExist(phone) { // 是否注册
		userId := userModel.GetUserIdForPwdAndAdmin(phone, pwd)
		if userId != "" { // 登录成功
			if userModel.IsAdmin(phone) { // 是否管理员
				resUserInfo := userModel.GetUserInfo(userId)
				service.SaveTokenRedis(userId, resUserInfo.Token) //保存到Redis
				return serializer.Response{
					Msg:  "登录成功",
					Data: resUserInfo,
				}
			} else {
				return serializer.Response{
					Code: 1,
					Msg:  "您暂未拥有管理员权限",
				}
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

// 获取用户信息
func UserInfo(userId string) serializer.Response {
	resUserInfo := userModel.GetUserInfo(userId)
	return serializer.Response{
		Code:0,
		Msg:  "获取成功",
		Data: resUserInfo,
	}
}

// 用户登出
func Logout(token string) serializer.Response {
	userId := userModel.GetUserIdForToken(token)
	if service.DeleteTokenRedis(userId) {
		return serializer.Response{
			Code: 0,
			Msg:  "用户已退出",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Msg:  "退出错误，请联系管理员",
		}
	}
}
