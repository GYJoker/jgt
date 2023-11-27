package files

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/GYJoker/jgt/cron"
	"io"
	"os"
	"time"
)

// TimerDeleteTempFile 定时删除临时文件
func TimerDeleteTempFile() {
	sts := &cron.YJCronTime{
		Job: &cron.YJJob{
			JobFunc: func() {
				deleteExpireTempFile(GetTempFileDir())
			},
		},
		TimeType: cron.JobTimeHour,
		Space:    2,
	}

	// 添加定时任务
	_, _ = sts.AddJob()
}

func deleteExpireTempFile(path string) {
	if !FileExists(path) {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	fileInfoList, err := file.Readdir(-1)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			deleteExpireTempFile(path + fileInfo.Name() + "/")
			continue
		}

		// 删除两小时以前的文件
		if fileInfo.ModTime().Unix() < (time.Now().Unix() - 7200) {
			RemoveFile(path + fileInfo.Name())
		}
	}
}

// FileExists 判断文件，文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// GetBaseDir 获取文件的基础目录
func GetBaseDir() string {
	return "./"
}

// GetConfigDir 获取配置文件的目录
func GetConfigDir() string {
	return GetBaseDir() + "./config.yaml"
}

// RemoveFile 删除文件
func RemoveFile(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
}

// GetTempFileDir 获取临时文件存储目录
func GetTempFileDir() string {
	return "temp/"
}

// GetSystemFilesDir 获取系统文件存储目录
func GetSystemFilesDir() string {
	return "files/"
}

// GetFileMd5 获取文件md5
func GetFileMd5(path string) string {
	open, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	hash := crypto.MD5.New()
	_, err = io.Copy(hash, open)
	if err != nil {
		fmt.Println(err)
	}
	err = open.Close()
	if err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
