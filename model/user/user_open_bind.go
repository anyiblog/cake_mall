package user

import (
	"cake_mall/conf"
	"cake_mall/util"
)

// UserOpenBind 第三方平台与用户绑定表
type UserOpenBind struct {
	OpenBindID string `gorm:"primary_key;column:open_bind_id;type:char(36);not null" json:"open_bind_id"` // 主键ID
	UserID     string `gorm:"column:user_id;type:char(36);not null" json:"user_id"`                       // 用户ID
	Openid     string `gorm:"column:openid;type:char(64);not null" json:"openid"`                         // 开放平台的ID
	OpenType   int    `gorm:"column:open_type;type:int(11);not null" json:"open_type"`                    // 开放平台类型
	OpenName   string `gorm:"column:open_name;type:varchar(20);not null" json:"open_name"`                // 开放平台名称
}

//新增用户开放平台绑定记录
func CreateOpenBind(userId,openId,openName string,openType int)  {
	openBindID := util.GenerateUUID()
	ob := UserOpenBind{
		OpenBindID: openBindID,
		UserID:     userId,
		Openid:     openId,
		OpenType:   openType,
		OpenName:   openName,
	}
	conf.DB.Create(ob)
}

////不存在第三方信息即绑定用户，自动注册
//func QueryIsExist(openInfo map[string]interface{}) bool {
//	fmt.Println(openInfo)
//	userOpenBind := &UserOpenBind{}
//	i := conf.DB.Where("openid = ?", openInfo["openid"]).Find(userOpenBind).RowsAffected
//	if i == 0 {
//		user := User{
//			UserID:     util.GenerateUUID(),
//			Phone:      openInfo["openid"].(string),
//			Password:   openInfo["openid"].(string),
//			CreateTime: time.Time{},
//		}
//		fmt.Println("user uuid", user.UserID)
//		conf.DB.Create(&user)
//		fmt.Println("创建用户成功")
//		imgId := model.UploadImg(openInfo["headimgurl"].(string))//图片表ID
//		userInfo := UserInfo{
//			UserInfoID:     util.GenerateUUID(),
//			UserID:         user.UserID,
//			Nickname:       openInfo["nickname"].(string),
//			Sex:            openInfo["sex"].(int),
//			Avatar:         imgId,
//			Balance:        0,
//			Integral:       0,
//			DefaultAddress: "",
//		}
//		conf.DB.Create(&userInfo)
//		fmt.Println("创建用户信息成功")
//	}
//	fmt.Println(i)
//	fmt.Println(userOpenBind)
//	return false
//}
