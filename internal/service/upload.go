/**
 * 文件上传-服务类
 * @since 2022/11/18
 * @File : upload
 */
package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/pkg/upload"
	"github.com/mwqnice/oh-admin/pkg/utils"
	"github.com/mwqnice/oh-admin/pkg/utils/gstr"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

// 上传得文件信息
type FileInfo struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	Src      string `json:"src"`
	FileType string `json:"fileType"`
}

func (svc *Service) UploadImage(ctx *gin.Context) (*FileInfo, error) {
	// 获取文件(注意这个地方的file要和html模板中的name一致)
	file, err := ctx.FormFile("file")
	if err != nil {
		return &FileInfo{}, errors.New("上传文件不能为空")
	}

	//获取文件的后缀名
	fileExt := path.Ext(file.Filename)

	// 允许上传文件后缀
	// 检查上传文件后缀
	if !upload.CheckContainExt(upload.TypeImage, file.Filename) {
		return &FileInfo{}, errors.New("上传文件格式不正确")
	}
	// 允许文件上传最大值
	allowSize := global.AppSetting.UploadImageMaxSize
	// 检查上传文件大小
	isvalid, err := upload.CheckFileSize(file.Size, allowSize)
	if err != nil {
		return &FileInfo{}, err
	}
	if !isvalid {
		return &FileInfo{}, errors.New("上传文件大小不得超过：" + allowSize)
	}

	// 存储目录
	savePath := utils.TempPath() + "/" + time.Now().Format("20060102")

	// 创建文件夹
	ok := utils.CreateDir(savePath)
	if !ok {
		return &FileInfo{}, errors.New("存储路径创建失败")
	}

	//根据当前时间鹾生成一个新的文件名
	fileNameInt := time.Now().Unix()
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	//新的文件名
	fileName := fileNameStr + fileExt
	//保存上传文件
	filePath := filepath.Join(savePath, "/", fileName)
	err2 := ctx.SaveUploadedFile(file, filePath)
	if err2 != nil {
		return &FileInfo{}, err2
	}
	// 返回结果
	result := &FileInfo{
		FileName: file.Filename,
		FileSize: file.Size,
		Src:      utils.GetImageUrl(gstr.Replace(savePath, utils.UploadPath(), "") + "/" + fileName),
	}
	return result, nil
}
