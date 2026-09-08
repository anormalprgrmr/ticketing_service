package handlers

import (
	"context"
	pb "ticket_service/internal/protos"
	"ticket_service/internal/services"
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
			UserID:  "",
		}, err
	}

	return &pb.NewUserResponse{
		Success: true,
		Error:   "",
		UserID:  userID.String(),
	}, nil
}
