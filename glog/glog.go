package glog

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"sync"
	"time"
)

type MyFormatter struct{}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

var (
	isInit         = false
	logPath string = "/logs"
)

func SetLogPath(path string) {
	logPath = path

	InitLog()
}

func InitLog() {
	writer, err := rotatelogs.New(
		logPath+"-%Y%m%d"+".log",
		rotatelogs.WithLinkName(logPath+".log"),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationCount(30),
		rotatelogs.WithRotationSize(10*1024*1024),

		// logFile+".%Y%m%d%H%M",                      //每分钟
		// rotatelogs.WithLinkName(logFile),           //生成软链，指向最新日志文件
		// rotatelogs.WithRotationTime(time.Minute),   //最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
		// rotatelogs.WithRotationCount(3),            //设置3份 大于3份 或到了清理时间 开始清理
		// rotatelogs.WithRotationSize(100*1024*1024), //设置100MB大小,当大于这个容量时，创建新的日志文件
	)
	if err != nil {
		panic(err)
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&MyFormatter{})
	logrus.SetOutput(writer)

	isInit = true
}

func GetLogger() *logrus.Logger {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	if !isInit {
		InitLog()
	}

	return logrus.StandardLogger()
}
