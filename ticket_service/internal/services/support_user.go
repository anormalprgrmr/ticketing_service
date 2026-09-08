package services

import (
	"context"
	"ticket_service/internal/models"
	"ticket_service/internal/repositories"
	"uuid"
)

type SupportService struct {
	supportRepo *repositories.SupportRepo
}

func NewSupportService(supportRepo *repositories.SupportRepo) *SupportService {
	return &SupportService{
		supportRepo: supportRepo,
	}
}

func (s *SupportService) CreateSupport(ctx context.Context, name string) (supportID uuid.UUID, err error) {

	supportID, err = s.supportRepo.CreateSupport(ctx, name)
	if err != nil {
		return uuid.Nil(), err
	}

	return supportID, err
}

func (s *SupportService) GetSupportTickets(ctx context.Context, name string) ([]models.Ticket, error) {

	tickets, err := s.supportRepo.GetSupportTickets(ctx, name)
	if err != nil {
		return nil, err
	}

	return tickets, err
}
