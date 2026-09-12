package services

import (
	"context"
	"ticket_service/internal/event"
	"ticket_service/internal/models"
	"ticket_service/internal/repositories"
	"uuid"
)

type TicketService struct {
	ticketRepo *repositories.TicketRepo
	eb         event.EventSignalBus
}

func NewTicketService(ticketRepo *repositories.TicketRepo, eb event.EventSignalBus) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
		eb:         eb,
	}
}

func (s *TicketService) CreateTicket(ctx context.Context, userID string, body string) (ticketID uuid.UUID, err error) {

	ticket, err := s.ticketRepo.NewTicket(ctx, userID, body)
	if err != nil {
		return uuid.Nil(), err
	}

	s.eb.Publish()

	return ticket.ID, err
}

func (s *TicketService) GetTicketsWithStatus(ctx context.Context, status models.TicketStatus) ([]*models.Ticket, error) {

	tickets, err := s.ticketRepo.GetTicketsWithStatus(ctx, string(status))
	if err != nil {
		return nil, err
	}

	return tickets, err
}

func (s *TicketService) CloseTicket(ctx context.Context, ticketId, supportID string) error {

	err := s.ticketRepo.CloseTicket(ctx, ticketId, supportID)
	if err != nil {
		return err
	}

	s.eb.Publish()

	return err
}

func (s *TicketService) AnswerTicket(ctx context.Context, ticketId, supportID, body string) error {

	err := s.ticketRepo.AnswerTicket(ctx, ticketId, supportID, body)
	if err != nil {
		return err
	}

	return err
}

func (s *TicketService) TransferTicket(ctx context.Context, ticketId, newSupportID string) error {

	err := s.ticketRepo.TransferTicket(ctx, ticketId, newSupportID)
	if err != nil {
		return err
	}

	return err
}
