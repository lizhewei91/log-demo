package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var log = logrus.New()

func init() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.Formatter = &logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	//log.SetOutput(os.Stdout)

	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("1.log", os.O_CREATE|os.O_WRONLY, 0666)
	writers := []io.Writer{
		file,
		os.Stdout,
	}

	// 同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.Infof("create log file faield, err:%v\n", err)
	}

	// 设置日志级别为warn以上
	log.Level = logrus.InfoLevel

	// 开启记录文件名和行号
	log.SetReportCaller(true)

	// 增加hook
	log.AddHook(NewDefaultFieldHook())
}

type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = "MyAppName"
	return nil
}
func (hook *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func NewDefaultFieldHook() *DefaultFieldHook {
	return &DefaultFieldHook{}
}

func main() {
	// 实例化一个logger对象
	contextLogger := log.WithFields(logrus.Fields{
		"component": "kce-scheduler-plugins",
	})

	contextLogger.Info("this is a info message")
	contextLogger.Debug("this is a debug message")
	contextLogger.Error("this is a error message")
}
