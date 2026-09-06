package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"ticket_service/internal/config"
	"ticket_service/internal/db"
	pb "ticket_service/internal/protos"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type TicketServer struct {
	dbConn *sql.DB
	pb.UnimplementedTicketServiceServer
}

func (s *TicketServer) NewTicket(ctx context.Context, in *pb.NewTicketRequest) (*pb.NewTicketResponse, error) {
	return &pb.NewTicketResponse{
		Value: "heeey",
	}, nil
}

func main() {

	err := config.InitLogger()

	config := config.NewConfig()

	dbConn := db.ConnectDB(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	dbConn.Query("SELECT * FROM tickets")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterTicketServiceServer(server, &TicketServer{})

	log.Info("Starting gRPC server ...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
