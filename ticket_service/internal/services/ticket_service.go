package services

import (
	"context"
	"ticket_service/internal/repositories"
	"uuid"
)

type TicketService struct {
	ticketRepo *repositories.TicketRepo
}

func NewTicketService(ticketRepo *repositories.TicketRepo) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
	}
}

func (s *TicketService) CreateTicket(ctx context.Context, userID string, body string) (ticketID uuid.UUID, err error) {

	ticket, err := s.ticketRepo.NewTicket(ctx, userID, body)
	if err != nil {
		return uuid.Nil(), err
	}

	return ticket.ID, err
}

