package service

import (
	"frozen-go-project/app/file_server/config"
	"frozen-go-project/app/file_server/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func UploadFile(uploadFile *multipart.FileHeader) (filepath, filename string, err error) {
	// 读取文件后缀
	ext := path.Ext(uploadFile.Filename)
	// 读取文件名并md5加密
	name := strings.TrimSuffix(uploadFile.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名(用时间戳)
	filename = name + "_" + time.Now().Format(utils.DATETIMEFORMAT1) + ext
	// 创建路径
	err = os.MkdirAll(config.UPLOAD_PATH, os.ModePerm)
	if err != nil {
		return
	}
	// 拼接路径和文件名
	filepath = config.UPLOAD_PATH + "/" + filename

	// 读取上传的文件
	f, err := uploadFile.Open()
	if err != nil {
		return
	}
	defer f.Close()

	// 创建文件
	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	// 拷贝文件
	_, err = io.Copy(out, f)
	if err != nil {
		return
	}
	return
}
