package handlers

import (
	"context"
	pb "ticket_service/internal/protos"
	"ticket_service/internal/services"
)

type TicketHandler struct {
	// pb.UnimplementedTicketServiceServer

	ticketService *services.TicketService
}

func NewTicketHandler(ticketService *services.TicketService) *TicketHandler {
	return &TicketHandler{
		ticketService: ticketService,
	}
}

func (h *TicketHandler) NewTicket(ctx context.Context, in *pb.NewTicketRequest) (*pb.NewTicketResponse, error) {
	ticketID, err := h.ticketService.CreateTicket(ctx, in.UserId, in.Body)
	if err != nil {
		return &pb.NewTicketResponse{
			Success:  false,
			Error:    err.Error(),
			TicketId: "",
		}, err
	}

	return &pb.NewTicketResponse{
		Success:  true,
		Error:    "",
		TicketId: ticketID.String(),
	}, nil
}

