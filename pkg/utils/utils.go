package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/utils/gmd5"
	"github.com/mwqnice/oh-admin/pkg/utils/gstr"
	"log"
	"math/rand"
	"os"
	"time"
)

//MD5 MD5加密
func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}

//EncodeMD5 加密
func EncodeMD5(str string) (string, error) {
	// 第一次MD5加密
	str, err := gmd5.Encrypt(str)
	if err != nil {
		return "", err
	}
	// 第二次MD5加密
	str, err = gmd5.Encrypt(str)
	if err != nil {
		return "", err
	}
	return str, nil
}

//GetClientIp 获取客户端IP
func GetClientIp(ctx *gin.Context) string {
	ip := ctx.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = ctx.ClientIP()
	}
	return ip
}

//InStringArray 是否在数组里
func InStringArray(url string, m map[string]interface{}) bool {
	_, ok := m[url]
	if ok {
		return true
	} else {
		return false
	}
}

// 判断元素是否在数组中
func InArray(value string, array []interface{}) bool {
	for _, v := range array {
		if convert.String(v) == value {
			return true
		}
	}
	return false
}

// 判断元素是否在数组中
func InIntArray(value int, array []int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// 附件目录
func UploadPath() string {
	// 获取项目根目录
	curDir, _ := os.Getwd()
	return curDir + "/" + global.AppSetting.UploadSavePath
}

// 临时目录
func TempPath() string {
	return UploadPath() + "/temp"
}

// 图片存放目录
func ImagePath() string {
	return UploadPath() + "/images"
}

// 文件目录(非图片目录)
func FilePath() string {
	return UploadPath() + "/file"
}

// 创建文件夹并设置权限
func CreateDir(path string) bool {
	// 判断文件夹是否存在
	if IsExist(path) {
		return true
	}
	// 创建多层级目录
	err2 := os.MkdirAll(path, os.ModePerm)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// 判断文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		if os.IsExist(err) {
			// 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}
func ImageUrl() string {
	return "/" + global.AppSetting.UploadSavePath
}

// 获取文件地址
func GetImageUrl(path string) string {
	return ImageUrl() + path
}

// GetRandomString 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	timeNano := time.Now().UnixNano()
	rand.Seed(timeNano + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

//GetRandomNumber 获取随机位数数字
func GetRandomNumber(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	timeNano := time.Now().UnixNano()
	rand.Seed(timeNano + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
func SaveImage(url string, dirname string) (string, error) {
	// 判断文件地址是否为空
	if gstr.Equal(url, "") {
		return "", errors.New("文件地址不能为空")
	}

	// 判断是否本站图片
	if gstr.Contains(url, ImageUrl()) {
		// 本站图片

		// 是否临时图片
		if gstr.Contains(url, "temp") {
			// 创建目录
			dirPath := ImagePath() + "/" + dirname + "/" + time.Now().Format("20060102")
			if !CreateDir(dirPath) {
				return "", errors.New("文件目录创建失败")
			}
			// 原始图片地址
			oldPath := gstr.Replace(url, ImageUrl(), UploadPath())
			// 目标目录地址
			newPath := ImagePath() + "/" + dirname + gstr.Replace(url, ImageUrl()+"/temp", "")
			// 移动文件
			os.Rename(oldPath, newPath)
			return GetImageUrl(gstr.Replace(newPath, UploadPath(), "")), nil
		} else {
			// 非临时图片
			//path := gstr.Replace(url, ImageUrl(), "")
			return url, nil
		}
	} else {
		// 远程图片
		// TODO...
	}
	return "", errors.New("保存文件异常")
}
