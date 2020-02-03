package user

import (
	"cake_mall/conf"
	"cake_mall/model"
	"cake_mall/util"
	"fmt"
	"os"
)

// UserInfo 用户信息表，保存用户余额，积分等数据
type UserInfo struct {
	UserInfoID     string  `gorm:"primary_key;column:user_info_id;type:char(36);not null" json:"user_info_id"`
	UserID         string  `gorm:"column:user_id;type:char(36);not null" json:"user_id"`                 // 用户id
	Nickname       string  `gorm:"column:nickname;type:char(20);not null" json:"nickname"`               // 昵称
	Sex            int     `gorm:"column:sex;type:int(11)" json:"sex"`                                   // 性别
	Avatar         string  `gorm:"column:avatar;type:char(36)" json:"avatar"`                            // 头像（图片表ID）
	Balance        float64 `gorm:"column:balance;type:decimal(18,2) unsigned;not null" json:"balance"`   // 余额
	Integral       int     `gorm:"column:integral;type:int(18) unsigned;not null" json:"integral"`       // 积分
	DefaultAddress string  `gorm:"column:default_address;type:char(36);not null" json:"default_address"` // 默认收货地址（地址表ID）
}

type ResUserInfo struct {
	Token          string `json:"token"`
	Phone          string `json:"phone"`
	Nickname       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Avatar         string `json:"avatar"`
	Balance        int    `json:"balance"`
	Integral       int    `json:"integral"`
	DefaultAddress string `json:"default_address"`
}

//创建普通用户信息
func CreateUserInfo(userId string) {
	var userInfo = struct {
		NickName  string `json:"nickName"`
		Gender    int    `json:"gender"`
		AvatarURL string `json:"avatarUrl"`
	}{
		"风铃用户_" + util.RandStr(5),
		0,
		os.Getenv("Default_User_Avatar"),
	}
	userInfoId := util.GenerateUUID()
	imgId := model.CreateImgForUrl(userInfo.AvatarURL)
	ui := UserInfo{
		UserInfoID:     userInfoId,
		UserID:         userId,
		Nickname:       userInfo.NickName,
		Sex:            userInfo.Gender,
		Avatar:         imgId,
		Balance:        0,
		Integral:       0,
		DefaultAddress: "",
	}
	conf.DB.Create(ui)
}

//创建微信用户信息
func CreateWxUserInfo(userId string, wxUserInfo util.WxUserInfo) {
	userInfoId := util.GenerateUUID()
	imgId := model.CreateImgForUrl(wxUserInfo.AvatarURL)
	ui := UserInfo{
		UserInfoID:     userInfoId,
		UserID:         userId,
		Nickname:       wxUserInfo.NickName,
		Sex:            wxUserInfo.Gender,
		Avatar:         imgId,
		Balance:        0,
		Integral:       0,
		DefaultAddress: "",
	}
	conf.DB.Create(ui)
}

//获取用户信息
func GetUserInfo(userId string) ResUserInfo {
	ResUserInfo := ResUserInfo{}
	//获取用户信息
	conf.DB.Model(UserInfo{}).Select("nickname,sex,avatar,balance,integral,default_address").Where("user_id = ?", userId).Scan(&ResUserInfo)
	fmt.Println(ResUserInfo)
	//获取token和手机
	ResUserInfo.Avatar = model.GetImgUrl(ResUserInfo.Avatar)
	token, phone := GetUser(userId)
	ResUserInfo.Token = token
	ResUserInfo.Phone = phone
	//获取地址
	ResUserInfo.DefaultAddress = GetDetailAddress(userId)
	return ResUserInfo
}
