package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("port is not provided")
	}

	return &Config{
		Port: portInt,
	}
}
