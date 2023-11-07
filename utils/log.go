package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	var (
		dirPath  = "./logs"
		fileName = "./logs/log"
		logLevel = "info"
	)

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	// To create logs folder
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// To create and write into log file inside log folder
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	Logger = &logrus.Logger{
		Out:   file,
		Level: level,
		Formatter: &logrus.JSONFormatter{
			// Time stamp in DD-MM-YYYY HH:MM:SS format
			TimestampFormat: "02-01-2006 15:04:05",
		},
	}

	Logger.Info("Logger started")
}

func GetLogger() *logrus.Logger {
	return Logger
}

func GetControllerLogger(controller string) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"controller": controller,
	})
}

func GetControllerLoggerWithFields(controller string, fields map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"controller": controller,
		"param":      fields,
	})
}

func GetFunctionLogger(function string) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"function": function,
	})
}

func GetFunctionLoggerWithFields(function string, fields map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"function": function,
		"param":    fields,
	})
}
