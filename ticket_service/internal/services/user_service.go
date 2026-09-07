package services

import (
	"context"
	"ticket_service/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepo
}

func NewUserService(userRepo *repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string) (userID string, err error) {

	user, err := s.userRepo.CreateUser(ctx, name)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
