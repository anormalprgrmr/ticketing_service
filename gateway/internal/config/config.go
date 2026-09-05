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
	if port == "" {
		port = "8081"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Println("wrong port number")
		portInt = 8081
	}

	cfg := Config{Port: portInt}

	return &cfg
}
