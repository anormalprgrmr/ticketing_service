package config

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(logLevel logrus.Level) error {
	logDir := "/tmp/newcash/gateway"

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	fileLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, "application.log"),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	log.SetOutput(io.MultiWriter(fileLogger, os.Stdout))
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetLevel(logLevel)
	log.Info("logger initialized successfully")

	return nil
}
