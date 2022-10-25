package golog

import (
	"io"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var Mlogger *logrus.Logger

func InitLogrusLog(writer io.Writer) {
	Mlogger = logrus.New()
	Mlogger.SetOutput(writer)
	Mlogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	Mlogger.SetReportCaller(true)
}

func InitLogrusRotateLog(path string) {
	writer, _ := rotatelogs.New(
		path + ".%Y%m%d",
	)
	Mlogger = logrus.New()
	Mlogger.SetOutput(writer)
	Mlogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	Mlogger.SetReportCaller(true)
}
