package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port       int
	RemoteHost string
	RemotePort int
}

func NewConfig() *Config {
	remoteHost := os.Getenv("TICKET-HOST")
	if remoteHost == "" {
		log.Fatal("TICKET-HOST is required")
	}

	remotePort := os.Getenv("TICKET-PORT")
	if remotePort == "" {
		log.Fatal("remotePort is required")
	}
	remotePortInt, err := strconv.Atoi(remotePort)
	if err != nil {
		log.Fatal("wrong port number")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Println("wrong port number")
		portInt = 8081
	}

	cfg := Config{
		Port:       portInt,
		RemoteHost: remoteHost,
		RemotePort: remotePortInt,
	}

	return &cfg
}
