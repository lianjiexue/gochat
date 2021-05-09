package utils

import (
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileRes struct {
	Code     uint
	Dir      string
	Name     string
	Dst      string
	Message  string
	SavePath string
}

//判断文件存储的目录
func TouchFilePath(file *multipart.FileHeader) FileRes {
	var res FileRes
	ext := filepath.Ext(file.Filename)
	if ext != ".png" && ext != ".jpeg" && ext != ".jpg" {
		res.Code = 0 //头像保存失败
		res.Message = "文件格式错误"
		return res
	}
	fileName := GetDateString() + ext
	dateString := time.Now().Format("2006/01/02")
	dirPath := "../upload/" + dateString

	if _, error := os.Stat(dirPath); os.IsNotExist(error) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		os.Chmod(dirPath, 0777)
		//如果文件夹不存在 则创建文件
	}
	dst := dirPath + "/" + fileName
	res.Code = 1
	res.Dst = dst
	res.Dir = dirPath
	res.Name = fileName
	res.SavePath = "/upload/" + dateString + "/" + fileName
	return res
}
