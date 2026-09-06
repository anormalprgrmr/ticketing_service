package config

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger

func InitLogger() *logrus.Logger {
	logger = logrus.New()
	fileLogger := &lumberjack.Logger{
		Filename:   "/tmp/newcash/gateway/application.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	logger.SetOutput(io.MultiWriter(fileLogger, os.Stdout))

	logger.Info("logger is initialized successfully!")

	return logger
}

func GetLogger() *logrus.Logger {
	return logger
}
