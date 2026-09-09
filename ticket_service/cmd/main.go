package main

import (
	"log"
	"ticket_service/internal/config"
	"ticket_service/internal/db"
	grpcserver "ticket_service/internal/grpc_server"
	ticketscheduler "ticket_service/internal/ticket_scheduler"
)

func main() {
	cfg := config.NewConfig()

	err := config.InitLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("cant init logger : %e", err)
	}

	dbConn := db.ConnectDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	ts := ticketscheduler.NewTicketScheduler(dbConn)

	err = grpcserver.StartgRPCServer(cfg.Port, dbConn, ts)

}
