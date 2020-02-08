package api

//根接口

import (
	"bytes"
	"cake_mall/model"
	"cake_mall/serializer"
	"cake_mall/service"
	"cake_mall/util"
	"cake_mall/util/sms"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

//服务器通达性
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Msg: "ok",
	})
}

//发送短信
func SendSms(c *gin.Context) {
	phone := c.Query("phone")
	smsType := c.Query("smsType")
	res := sms.SendSms(phone, smsType)
	c.JSON(200, res)
}

//根据图片链接删除图片
func DeleteFile(c *gin.Context) {
	filePath := c.PostForm("imgUrl")                                    //filePath带域名，删除数据库需要
	fileName := strings.Replace(filePath, os.Getenv("COS_Url"), "", -1) //fileName不带，删除Cos需要
	cosErr := service.DeleteCosFile(fileName)                           //删除Cos文件
	if cosErr == nil && model.DeleteImg(filePath) == true {
		c.JSON(200, serializer.Response{
			Code: 0,
			Msg:  "删除文件成功",
		})
	} else {
		if !model.DeleteImg(filePath) {
			c.JSON(200, serializer.Response{
				Code: 1,
				Msg:  "删除文件失败，服务器错误。",
			})
		}
		if cosErr != nil {
			c.JSON(200, serializer.Response{
				Code: 1,
				Msg:  cosErr.Error(),
			})
		}
	}
}

//文件上传
func UploadFile(c *gin.Context) {
	fmt.Println(c.ContentType())
	if c.ContentType() == "application/json" { //json 批量上传
		type UploadFileList struct {
			List []string `json:"list" binding:"required"`
		}
		var uploadFileList UploadFileList
		if err := c.ShouldBind(&uploadFileList); err == nil {
			paramUrlList := uploadFileList.List //图片链接列表
			var resFileList []string            //上传成功后CosUrl文件列表
			var uploadOkFileList []string       //已上传成功的文件名
			for i := range paramUrlList {
				fmt.Println(paramUrlList[i])
				//fileName （uuid+文件名后缀） filePath （本地存储路径 ./tmp/....png）
				fileName, filePath, downloadRes := getImg(paramUrlList[i])
				if downloadRes.Code == 1 { //如果文件类型不是图片类型，删除Cos已上传的图片
					c.JSON(200, downloadRes)
					for i := range uploadOkFileList {
						err := service.DeleteCosFile(uploadOkFileList[i])
						fmt.Println(err)
					}
					return
				} else {
					fDataByte, _ := ioutil.ReadFile(filePath)   //byte字节
					fDataIoReader := bytes.NewReader(fDataByte) //byte字节转io.reader类型
					res := uploadFile(fDataIoReader, fileName)
					resFileList = append(resFileList, res.Data.(string))
					uploadOkFileList = append(uploadOkFileList, fileName)
					_ = os.Remove(filePath) //删除本地缓存文件
				}
			}
			c.JSON(200, serializer.Response{ //将图片记录插入数据库
				Code: 0,
				Msg:  "上传成功",
				Data: resFileList,
			})
			for i := range resFileList {
				model.CreateImg(resFileList[i])
			}
		} else {
			c.JSON(200, serializer.Response{
				Code: 1,
				Msg:  "参数错误",
			})
		}
	} else if c.ContentType() == "multipart/form-data" { //form-data 单文件上传
		form, _ := c.MultipartForm()
		// 获取所有图片
		files := form.File["file"]
		var resFileList []string      //上传成功后CosUrl文件列表
		var uploadOkFileList []string //已上传成功的文件名
		for _, file := range files {
			fData, _ := file.Open() //io.reader
			defer fData.Close()
			fileExt := path.Ext(file.Filename) //获取文件后缀
			fileName := util.GenerateUUID() + fileExt
			res := uploadFile(fData, fileName)
			if res.Code == 0 {
				resFileList = append(resFileList, res.Data.(string))
				uploadOkFileList = append(uploadOkFileList, fileName)
			} else {
				c.JSON(200, res)
				for i := range uploadOkFileList {
					err := service.DeleteCosFile(uploadOkFileList[i])
					fmt.Println(err)
				}
				return
			}
		}
		c.JSON(200, serializer.Response{
			Code: 0,
			Msg:  "文件上传成功",
			Data: resFileList,
		})
		for i := range resFileList { //将图片记录插入数据库
			model.CreateImg(resFileList[i])
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "请求类型错误",
		})
	}
}

//文件上传具体实现（二进制流，文件名）
func uploadFile(fData io.Reader, fileName string) serializer.Response {
	fileUrl, fileErr := service.CosUpload(fData, fileName)
	//fmt.Println(fileUrl)
	if fileErr != nil {
		res := serializer.Response{
			Code: 1,
			Msg:  fileErr.Error(),
		}
		return res
	} else {
		res := serializer.Response{
			Code: 0,
			Msg:  "文件上传成功",
			Data: fileUrl,
		}
		return res
	}
}

//保存文件
func getImg(url string) (fileName, filePath string, res serializer.Response) {
	imagPath := url
	resp, _ := http.Get(imagPath)
	imgType := resp.Header.Get("Content-Type")
	if imgType == "image/png" || imgType == "image/jpeg" || imgType == "image/gif" {
		a := strings.Split(imgType, "/")
		body, _ := ioutil.ReadAll(resp.Body)
		savePath := os.Getenv("Tmp_Path")
		fileName = util.GenerateUUID() + "." + a[1]
		filePath = savePath + fileName
		//通过http请求获取图片的流文件
		out, _ := os.Create(filePath)
		_, _ = io.Copy(out, bytes.NewReader(body))
		res := serializer.Response{
			Code: 0,
		}
		return fileName, filePath, res
	} else {
		res := serializer.Response{
			Code: 1,
			Msg:  "URL不是正确的图片链接",
		}
		return "", "", res
	}
}
