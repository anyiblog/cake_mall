//工具公用函数库

package util

import (
	md52 "crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//随机字符串种子
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

//生成随机字符串
func RandStr(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//生成Token
func GenerateToken() string {
	return RandStr(28) + "+" + RandStr(8) + "+" + RandStr(26)
}

//生成UUID
func GenerateUUID() string {
	return uuid.New().String()
}

//生成MD5
func Md5(str string) string {
	data := []byte(str)
	md5Ctx := md52.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	strMd5 := fmt.Sprintf("%x", cipherStr)
	return strMd5
}

//字符串JSON转map
func StrJsonToMap(strJson string) map[string]string {
	var strMap map[string]string
	if err := json.Unmarshal([]byte(strJson), &strMap); err != nil {
		fmt.Println(err)
	}
	return strMap
}

//返回多少天后的秒数，Redis缓存会用
func MinToSm(day string) int {
	now := time.Now()
	ad, _ := time.ParseDuration("24h")
	fmt.Println(now.Add(ad * 7).Unix())
	i := now.Add(ad*7).Unix() - now.Unix()
	return int(i)
}

//返回多少天后的秒数，Redis缓存会用
func DayToSm(day string) int {
	now := time.Now()
	ad, _ := time.ParseDuration("24h")
	fmt.Println(now.Add(ad * 7).Unix())
	i := now.Add(ad*7).Unix() - now.Unix()
	return int(i)
}

//结构体转json
func StructToJson(structData interface{}) string {
	jsons, errs := json.Marshal(structData) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	return string(jsons)
}

//验证手机号
func CheckPhone(phone string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

func SixRandNum() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

func TimeFormat(timeT time.Time) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := timeT.Format(timeLayout)
	loc, _ := time.LoadLocation("Local")                         //时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //字符串转换Time类型
	return theTime
}

func NowTime() time.Time {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Now().Format(timeLayout)
	loc, _ := time.LoadLocation("Local")                         //时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //字符串转换Time类型
	return theTime
}
