package services

import (
	"context"
	"ticket_service/internal/repositories"
)

type SupportService struct {
	supportRepo *repositories.SupportRepo
}

func NewSupportService(supportRepo *repositories.SupportRepo) *SupportService {
	return &SupportService{
		supportRepo: supportRepo,
	}
}

func (s *SupportService) CreateSupport(ctx context.Context, name string) (supportID string, err error) {

	supportID, err = s.supportRepo.CreateSupport(ctx, name)

	return supportID, err
}
