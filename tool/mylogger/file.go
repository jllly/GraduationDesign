package mylogger

import (
	"fmt"
	"github.com/go-playground/locales/fur"
	"time"
)

//往文件里面写日志相关代码

type LogLevel string

type FileLogger struct {
	Lever       LogLevel //日志的级别
	filePath    string   //日志文件保存的路径
	fileName    string   //日志文件保存的文件名
	maxFileSize int64    //按时间切割（还没修改）
}

//NewFileLogger构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	return &FileLogger{
		Lever:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= l.level
}
func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprint(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), file)
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		log(INFO, format, a...)
	}
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.enable(WARNING) {
		log(WARNING, format, a...)
	}
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		log(ERROR, format, a...)
	}
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAL) {
		log(FATAL, format, a...)
	}
}
