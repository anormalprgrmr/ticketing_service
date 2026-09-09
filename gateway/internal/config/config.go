package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port       int
	RemoteHost string
	RemotePort int
	LogLevel   log.Level
}

func NewConfig() (*Config, error) {
	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	remoteHost := os.Getenv("TICKET-HOST")
	if remoteHost == "" {
		log.Fatal("TICKET-HOST is required")
	}

	remotePort := os.Getenv("TICKET-PORT")
	if remotePort == "" {
		log.Fatal("TICKET-PORT is required")
	}
	remotePortInt, err := strconv.Atoi(remotePort)
	if err != nil {
		log.Fatal("wrong TICKET-PORT type")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Debugf("port is not specified, using default port 8081")
		port = "8081"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Println("wrong PORT type")
		portInt = 8081
	}

	cfg := Config{
		Port:       portInt,
		RemoteHost: remoteHost,
		RemotePort: remotePortInt,
		LogLevel:   logLevel,
	}
	return &cfg, nil
}
