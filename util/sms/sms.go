package sms

import (
	"cake_mall/cache"
	"cake_mall/serializer"
	"cake_mall/util"
	"fmt"
	"github.com/sasaxie/qcloundsms-go/sms"
	"os"
	"strconv"
	"time"
)

func SendSms(phone string, smsType string) serializer.Response {
	SmsAppId, _ := strconv.Atoi(os.Getenv("SMS_APP_ID"))
	SmsAppKey := os.Getenv("SMS_APP_KEY")

	smsCode := util.SixRandNum()
	cache.NewCache(util.RedisSms, "sms_").Set(phone+"_"+smsType, smsCode, 5*time.Minute)

	// 短信模版内容
	params := make([]string, 0)
	params = append(params, smsCode)
	params = append(params, "5")

	// 短信模版ID，需要在短信应用中申请
	templateId := smsType
	templateIdInt, _ := strconv.Atoi(templateId) //转int
	// 签名参数使用的是`签名内容`，而不是`签名ID`
	smsSign := "安忆博客"
	fmt.Println(SmsAppId)
	singleSender := sms.NewSingleSender(SmsAppId, SmsAppKey)
	fmt.Println(util.SixRandNum())
	result, _ := singleSender.SendWithParam("86", phone, templateIdInt, params, smsSign, "", "")
	return serializer.Response{
		Code: result.Result,
		Data: nil,
		Msg:  result.ErrMsg,
	}
}
