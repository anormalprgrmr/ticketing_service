package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"ticket_service/config"
	pb "ticket_service/internal/protos"

	"google.golang.org/grpc"
)

type TicketServer struct {
	pb.UnimplementedTicketServiceServer
}

func (s *TicketServer) NewTicket(ctx context.Context, in *pb.NewTicketRequest) (*pb.NewTicketResponse, error) {
	return &pb.NewTicketResponse{
		Value: "heeey",
	}, nil
}

func main() {
	config := config.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterTicketServiceServer(server, &TicketServer{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
