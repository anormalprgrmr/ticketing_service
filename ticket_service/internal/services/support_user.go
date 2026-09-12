package services

import (
	"context"
	"ticket_service/internal/event"
	"ticket_service/internal/models"
	"ticket_service/internal/repositories"
	"uuid"
)

type SupportService struct {
	supportRepo *repositories.SupportRepo
	eb          event.EventSignalBus
}

func NewSupportService(supportRepo *repositories.SupportRepo, eb event.EventSignalBus) *SupportService {
	return &SupportService{
		supportRepo: supportRepo,
		eb:          eb,
	}
}

func (s *SupportService) CreateSupport(ctx context.Context, name string) (supportID uuid.UUID, err error) {

	supportID, err = s.supportRepo.CreateSupport(ctx, name)
	if err != nil {
		return uuid.Nil(), err
	}

	s.eb.Publish()

	return supportID, err
}

func (s *SupportService) GetSupportTickets(ctx context.Context, name string, pageSize, pageNum int64) ([]*models.Ticket, error) {

	tickets, err := s.supportRepo.GetSupportTickets(ctx, name, pageSize, pageNum)
	if err != nil {
		return nil, err
	}

	return tickets, err
}
