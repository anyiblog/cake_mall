package user

import (
	"cake_mall/conf"
	"cake_mall/util"
	"time"
)

// User 用户表
type User struct {
	UserID    string    `gorm:"primary_key;column:user_id;type:char(36);not null" json:"user_id"`
	Phone     string    `gorm:"column:phone;type:char(11);not null" json:"phone"`       // 手机号
	Password  string    `gorm:"column:password;type:char(32);not null" json:"password"` // 密码
	Token     string    `gorm:"column:token;type:char(64);not null" json:"token"`       // 用户token
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`     // 注册时间
	Admin     int       `gorm:"column:admin;type:int(10) unsigned" json:"admin"`        // 是否为管理员
}

//刷新Token（每次登录会刷新）
func RefreshToken(userId, token string) {
	conf.DB.Model(User{}).Where("user_id = ?", userId).Update(map[string]interface{}{"token": token})
}

// 是否存在用户 true 存在 false 不存在
func IsExist(phone string) bool {
	i := 0
	conf.DB.Model(User{}).Where("phone = ?", phone).Count(&i)
	if i > 0 {
		return true
	} else {
		return false
	}
}

// 是否是管理员 true 是管理员 false 不是
func IsAdmin(phone string) bool {
	var admin struct {
		Admin int // 0 为拥有管理员权限，大于 0 即拥有管理权限 （埋坑，后续可能会有权限系统）
	}
	conf.DB.Model(User{}).Select("admin").Where("phone = ?", phone).Scan(&admin)
	if admin.Admin > 0 {
		return true
	} else {
		return false
	}
}

// 是否是管理员 true 是管理员 false 不是
func IsAdminForToken(token string) bool {
	var admin struct {
		Admin int // 0 为拥有管理员权限，大于 0 即拥有管理权限 （埋坑，后续可能会有权限系统）
	}
	conf.DB.Model(User{}).Select("admin").Where("token = ?", token).Scan(&admin)
	if admin.Admin > 0 {
		return true
	} else {
		return false
	}
}

// 创建普通用户 return userId
func CreateUser(phone, pwd string) string {
	userId := util.GenerateUUID()
	token := util.GenerateToken()
	u := User{
		UserID:    userId,
		Phone:     phone,
		Password:  util.Md5(pwd),
		Token:     token,
		CreatedAt: util.NowTime(),
	}
	conf.DB.Create(u) //创建用户
	CreateUserInfo(userId)
	return userId
}

// 有记录返回用户信息，无则创建微信用户 return userId
func CreateWxUser(wxUserInfo util.WxUserInfo) string {
	i := conf.DB.Where("openid = ?", wxUserInfo.OpenID).Take(&UserOpenBind{})
	if i.RowsAffected == 0 {
		userId := util.GenerateUUID()
		token := util.GenerateToken()
		//拼接手机号为：开放平台名称+下划线+随机字符串
		openName := "WeChat"
		strSum := 11 - len(openName) - 1 //11为手机号位数 1为下划线
		strPhone := "WeChat" + "_" + util.RandStr(strSum)
		u := User{
			UserID:    userId,
			Phone:     strPhone,
			Password:  util.RandStr(32),
			Token:     token,
			CreatedAt: util.NowTime(),
		}
		conf.DB.Create(u) //创建用户

		CreateWxUserInfo(u.UserID, wxUserInfo)                                     //创建用户信息
		CreateOpenBind(u.UserID, wxUserInfo.OpenID, openName, util.OpenTypeWeChat) //绑定用户第三方openID
		return u.UserID
	} else {
		var userId = struct {
			UserId string
		}{}
		i.Select("user_id").Scan(&userId)
		return userId.UserId
	}
}

// 根据token返回UserId
func GetUserIdForToken(token string) string {
	var userId struct {
		UserId string
	}
	conf.DB.Model(User{}).Select("user_id").Where("token = ? ", token).Scan(&userId)
	return userId.UserId
}

// 手机+密码获取userId 并且必须是管理员
func GetUserIdForPwdAndAdmin(phone, pwd string) string {
	var userId struct {
		UserId string
	}
	conf.DB.Model(User{}).Select("user_id").Where("phone = ? and password = ? ", phone, util.Md5(pwd)).Scan(&userId)
	return userId.UserId
}

//手机+密码获取userId
func GetUserIdForPwd(phone, pwd string) string {
	var userId struct {
		UserId string
	}
	conf.DB.Model(User{}).Select("user_id").Where("phone = ? and password = ?", phone, util.Md5(pwd)).Scan(&userId)
	return userId.UserId
}

//手机获取userId
func GetUserIdForPhone(phone string) string {
	var userId struct {
		UserId string
	}
	conf.DB.Model(User{}).Select("user_id").Where("phone = ?", phone).Scan(&userId)
	return userId.UserId
}

//获取用户token和手机号
func GetUser(userId string) (token, phone string) {
	var user struct {
		Token string
		Phone string
	}
	conf.DB.Model(User{}).Select("phone,token").Where("user_id = ?", userId).Scan(&user)
	return user.Token, user.Phone
}

//如果手机号是随机字符串则绑定新手机号（仅适用第三方平台无手机号）
func BindPhoneAndPwd(userId, phone, pwd string) bool {
	if conf.DB.Model(User{}).Where("user_id = ?", userId).Update(map[string]interface{}{"phone": phone, "password": util.Md5(pwd)}).RecordNotFound() {
		return false
	} else {
		return true
	}
}

//通过验证码更新密码
func UpdatePwd(phone, newPwd string) bool {
	if conf.DB.Model(User{}).Where("phone = ?", phone).Update(map[string]interface{}{"password": util.Md5(newPwd)}).RecordNotFound() {
		return false
	} else {
		return true
	}
}
