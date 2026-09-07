package config

import (
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel   logrus.Level
	Port       int
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func NewConfig() *Config {

	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("port is not provided")
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	dbPortInt, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalf("POSTGRES_PORT is not provided")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		log.Fatalf("POSTGRES_HOST is not provided")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		log.Fatalf("POSTGRES_USER is not provided")
	}

	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		log.Fatalf("POSTGRES_PASSWORD is not provided")
	}

	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		log.Fatalf("POSTGRES_DB is not provided")
	}
	return &Config{
		Port:       portInt,
		DBPort:     dbPortInt,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		LogLevel:   logLevel,
	}
}
