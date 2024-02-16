package files

import (
	"archive/zip"
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/GYJoker/jgt/cron"
	"io"
	"os"
	"path/filepath"
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

// ZipFileDir 压缩文件
func ZipFileDir(path string) (string, error) {
	// 获取输入路径的目录部分
	dirPath := filepath.Dir(path)

	// 基于输入路径生成ZIP文件名
	zipFileName := filepath.Join(dirPath, fmt.Sprintf("%s.zip", filepath.Base(path)))

	// 创建ZIP文件
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return "", err
	}
	defer func(zipFile *os.File) {
		e := zipFile.Close()
		if e != nil {

		}
	}(zipFile)

	// 创建ZIP写入器
	zw := zip.NewWriter(zipFile)
	defer func(zw *zip.Writer) {
		er := zw.Close()
		if er != nil {

		}
	}(zw)

	// 递归函数，用于将文件夹或文件添加到ZIP中
	var addToZip func(string, string) error
	addToZip = func(srcPath, zipPath string) error {
		fileInfo, e := os.Stat(srcPath)
		if e != nil {
			return e
		}

		if fileInfo.IsDir() {
			// 如果是文件夹，递归添加文件夹内的文件
			entries, er := os.ReadDir(srcPath)
			if er != nil {
				return er
			}
			for _, entry := range entries {
				if r := addToZip(filepath.Join(srcPath, entry.Name()), filepath.Join(zipPath, entry.Name())); err != nil {
					return r
				}
			}
		} else {
			// 如果是文件，将其添加到ZIP中
			file, er := os.Open(srcPath)
			if er != nil {
				return er
			}
			defer func(file *os.File) {
				r := file.Close()
				if r != nil {

				}
			}(file)

			// 创建ZIP文件条目
			f, r := zw.Create(zipPath)
			if r != nil {
				return r
			}

			// 将文件内容写入ZIP条目
			_, ee := io.Copy(f, file)
			return ee
		}
		return nil
	}

	// 判断输入路径是文件还是文件夹，并添加到ZIP中
	if er := addToZip(path, filepath.Base(path)); er != nil {
		return "", er
	}

	return zipFileName, nil
}
