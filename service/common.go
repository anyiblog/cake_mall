//业务逻辑公用函数库

package service

import (
	"cake_mall/cache"
	"cake_mall/util"
	"context"
	"github.com/nelsonken/cos-go-sdk-v5/cos"
	"io"
	"os"
)

//简单 Cos上传文件
func CosUpload(content io.Reader, path string) (fileUrl string, errInfo error) {
	cc := cos.New(&cos.Option{
		AppID:     os.Getenv("APP_ID"),
		SecretID:  os.Getenv("SECRET_ID"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Region:    os.Getenv("COS_REGION"),
	})
	errInfo = cc.Bucket(os.Getenv("COS_BUCKET_NAME")).UploadObject(context.Background(), path, content, &cos.AccessControl{})
	fileUrl = os.Getenv("COS_Url") + path
	return fileUrl, errInfo
}

//删除Cos文件
func DeleteCosFile(fileName string) (errInfo error) {
	cc := cos.New(&cos.Option{
		AppID:     os.Getenv("APP_ID"),
		SecretID:  os.Getenv("SECRET_ID"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Region:    os.Getenv("COS_REGION"),
	})
	errInfo = cc.Bucket(os.Getenv("COS_BUCKET_NAME")).DeleteObject(context.Background(), fileName)
	return errInfo
}

//检测短信验证码Redis值是否相等
func CheckSmsRedis(phone, SmsType, smsCode string) bool {
	c := cache.NewCache(util.RedisSms, "sms_")
	key := phone + "_" + SmsType
	value := c.Get(key)
	if value == smsCode {
		return true
	} else {
		return false
	}
}
