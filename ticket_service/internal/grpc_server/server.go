package grpcserver

import (
	"database/sql"
	"fmt"
	"net"
	"ticket_service/internal/handlers"
	pb "ticket_service/internal/protos"
	"ticket_service/internal/repositories"
	"ticket_service/internal/services"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type TicketServer struct {
	dbConn *sql.DB
	*handlers.TicketHandler
	*handlers.UserHandler
}

func StartgRPCServer(port int, dbConn *sqlx.DB) error {

	ticketRepo := repositories.NewTicketRepo(dbConn)
	userRepo := repositories.NewUserRepo(dbConn)

	ticketService := services.NewTicketService(ticketRepo)
	userService := services.NewUserService(userRepo)

	ticketHandler := handlers.NewTicketHandler(ticketService)
	userHandler := handlers.NewUserHandler(userService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterTicketServiceServer(server, &TicketServer{
		TicketHandler: ticketHandler,
		UserHandler:   userHandler,
	})

	log.Info("Starting gRPC server ...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
