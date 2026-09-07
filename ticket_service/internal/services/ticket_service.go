package services

import "ticket_service/internal/repositories"

type TicketService struct {
	ticketRepo *repositories.TicketRepo
}

func NewTicketService(ticketRepo *repositories.TicketRepo) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
	}
}

func (s *TicketService) CreateTicket(name string) error {

	s.ticketRepo.NewTicket("sad", "ss")

	return nil
}
