package services

import "ticket_service/internal/repositories"

type SupportService struct {
	supportRepo *repositories.SupportRepo
}

func NewSupportService(supportRepo *repositories.SupportRepo) *SupportService {
	return &SupportService{
		supportRepo: supportRepo,
	}
}

func (s *SupportService) CreateSupport(name string) error {

	s.supportRepo.CreateSupport(name)

	return nil
}
