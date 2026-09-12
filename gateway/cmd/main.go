package main

import (
	"fmt"
	"gateway/internal/config"
	pb "gateway/internal/protos"
	"gateway/internal/routers"
	"net/http"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Info("Starting gateway App...")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("cant load config: %e", err)
	}

	err = config.InitLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("cant create logger instance : %e", err)
	}

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.RemoteHost, cfg.RemotePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cant connect to gRPC server :%v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	r := routers.InitRouter(client)

	log.Printf("🫸🫸🫸 Starting HTTP server... Using port %d", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		log.Fatalf("error starting HTTP gateway %e", err)
	}
}
