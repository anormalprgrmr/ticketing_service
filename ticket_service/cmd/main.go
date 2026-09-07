package main

import (
	"log"
	"ticket_service/internal/config"
	"ticket_service/internal/db"
	grpcserver "ticket_service/internal/grpc_server"
)

func main() {

	err := config.InitLogger()
	if err != nil {
		log.Fatalf("cant init logger : %e", err)
	}

	config := config.NewConfig()

	dbConn := db.ConnectDB(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	err = grpcserver.StartgRPCServer(config.Port, dbConn)

}
