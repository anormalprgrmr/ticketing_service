package services

import (
	"context"
	"ticket_service/internal/models"
	"ticket_service/internal/repositories"
	"uuid"
)

type UserService struct {
	userRepo *repositories.UserRepo
}

func NewUserService(userRepo *repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string) (userID uuid.UUID, err error) {

	user, err := s.userRepo.CreateUser(ctx, name)
	if err != nil {
		return uuid.Nil(), err
	}

	return user.ID, nil
}

func (s *UserService) GetUserTickets(ctx context.Context, userID string) ([]*models.Ticket, error) {

	tickets, err := s.userRepo.GetUserTickets(ctx, userID)
	if err != nil {
		return nil, err
	}

	return tickets, err
}
