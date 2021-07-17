// author pengchengbai@shopee.com
// date 2021/7/17

package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)
import "github.com/rifflock/lfshook"

var log *logrus.Logger
var writer io.WriteCloser

func NewLoggerWithRotate() *logrus.Logger {
	if log != nil {
		return log
	}

	path := "./log/info.log"
	writer, _ = rotatelogs.New(
		path+".%Y%m%d.%H%M",
		rotatelogs.WithLinkName(path),               // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(60*time.Second),       // 文件最大保存时间
		rotatelogs.WithRotationTime(20*time.Second), // 日志切割时间间隔
	)

	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.PanicLevel: writer,
	}

	log = logrus.New()
	// 用于打印日志所在文件的位置, 参考：https://blog.csdn.net/wslyk606/article/details/81670713
	log.AddHook(lineHook{})
	log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.TextFormatter{},
	))

	return log
}

func Close()  {
	if writer != nil {
		writer.Close()
	}
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args)
}

func Error(args ...interface{}) {
	log.Error(args)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

func Panic(args ...interface{}) {
	log.Panic(args)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args)
}
