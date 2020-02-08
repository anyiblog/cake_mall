package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
)

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

type WxUserInfo struct {
	OpenID    string `json:"openid"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Watermark struct {
		Timestamp int    `json:"timestamp"`
		Appid     string `json:"appid"`
	} `json:"watermark"`
}

func WxDataDecrypt(aesKey, encryptedData, iv string) WxUserInfo {
	var wxu WxUserInfo
	ak, _ := base64.StdEncoding.DecodeString(aesKey)
	i, _ := base64.StdEncoding.DecodeString(iv)
	data, _ := base64.StdEncoding.DecodeString(encryptedData)
	aesStr, _ := AesDecrypt(data, ak, i)
	_ = json.Unmarshal([]byte(aesStr), &wxu)
	return wxu
}
