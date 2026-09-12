package handlers

import (
	"context"
	pb "ticket_service/internal/protos"
	"ticket_service/internal/services"
	"ticket_service/internal/transformers"
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

func (h *TicketHandler) GetTicketsWithStatus(ctx context.Context, in *pb.GetTicketsWithStatusRequest) (*pb.GetTicketsWithStatusResponse, error) {
	tickets, err := h.ticketService.GetTicketsWithStatus(ctx, transformers.ConvertTicketStatusGRPCToModel(in.Status))
	if err != nil {
		return &pb.GetTicketsWithStatusResponse{
			Success: false,
			Error:   err.Error(),
			Tickets: nil,
		}, err
	}

	return &pb.GetTicketsWithStatusResponse{
		Success: true,
		Error:   "",
		Tickets: transformers.TicketModelToGRPC(tickets),
	}, nil
}

func (h *TicketHandler) CloseTicket(ctx context.Context, in *pb.CloseTicketRequest) (*pb.CloseTicketResponse, error) {
	err := h.ticketService.CloseTicket(ctx, in.TicketId, in.SupportId)
	if err != nil {
		return &pb.CloseTicketResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &pb.CloseTicketResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (h *TicketHandler) AnswerTicket(ctx context.Context, in *pb.AnswerTicketRequest) (*pb.AnswerTicketResponse, error) {
	err := h.ticketService.AnswerTicket(ctx, in.TicketId, in.SupportId, in.Body)
	if err != nil {
		return &pb.AnswerTicketResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &pb.AnswerTicketResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (h *TicketHandler) TransferTicket(ctx context.Context, in *pb.TransferTicketRequest) (*pb.TransferTicketResponse, error) {
	err := h.ticketService.TransferTicket(ctx, in.TicketId, in.SupportId)
	if err != nil {
		return &pb.TransferTicketResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &pb.TransferTicketResponse{
		Success: true,
		Error:   "",
	}, nil
}
