package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"ticket_service/internal/config"
	"ticket_service/internal/db"
	"ticket_service/internal/event"
	grpcserver "ticket_service/internal/grpc_server"
	ticketscheduler "ticket_service/internal/ticket_scheduler"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting ticket App...")

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	cfg := config.NewConfig()

	err := config.InitLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("cant init logger : %e", err)
	}

	dbConn := db.ConnectDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	eb := event.NewChannelSignalBus()

	ts := ticketscheduler.NewTicketScheduler(dbConn, eb)

	err = ts.Start(ctx)
	if err != nil {
		log.Fatalf("couldnt start tickerScheduler : %e", err)
	}

	err = grpcserver.StartgRPCServer(cfg.Port, dbConn, eb)

}
