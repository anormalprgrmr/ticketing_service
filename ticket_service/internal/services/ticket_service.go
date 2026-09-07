package services

import (
	"context"
	"ticket_service/internal/repositories"
)

type TicketService struct {
	ticketRepo *repositories.TicketRepo
}

func NewTicketService(ticketRepo *repositories.TicketRepo) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
	}
}

func (s *TicketService) CreateTicket(ctx context.Context, userID string, body string) (ticketID string, err error) {

	ticket, err := s.ticketRepo.NewTicket(ctx, userID, body)

	return ticket.ID, err
}
