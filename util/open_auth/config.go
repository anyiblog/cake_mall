package open_auth

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	Wechatappkey    = ""
	Wechatappsecret = ""

	QQappkey      = os.Getenv("QQ_APPKEY")
	QQappsecret   = os.Getenv("QQ_APPSECRET")
	QQRedirectUrl = os.Getenv("QQ_REDIRECT_URL")

	Weiboappkey      = ""
	Weiboappsecret   = ""
	WeiboRedirectUrl = ""
)
