package handlers

import (
	"context"
	pb "ticket_service/internal/protos"
	"ticket_service/internal/services"
	"ticket_service/internal/transformers"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) NewUser(ctx context.Context, in *pb.NewUserRequest) (*pb.NewUserResponse, error) {

	userID, err := h.userService.CreateUser(ctx, in.Name)
	if err != nil {
		return &pb.NewUserResponse{
			Success: false,
			Error:   err.Error(),
			UserId:  "",
		}, err
	}

	return &pb.NewUserResponse{
		Success: true,
		Error:   "",
		UserId:  userID.String(),
	}, nil
}

func (h *UserHandler) GetUserTickets(ctx context.Context, in *pb.GetUserTicketsRequest) (*pb.GetUserTicketsResponse, error) {

	tickets, err := h.userService.GetUserTickets(ctx, in.UserId)
	if err != nil {
		return &pb.GetUserTicketsResponse{
			Success: false,
			Error:   err.Error(),
			Tickets: nil,
		}, err
	}

	return &pb.GetUserTicketsResponse{
		Success: true,
		Error:   "",
		Tickets: transformers.TicketModelToGRPC(tickets),
	}, nil
}
