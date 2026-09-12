package handlers

import (
	"context"
	pb "ticket_service/internal/protos"
	"ticket_service/internal/services"
	"ticket_service/internal/transformers"

	"github.com/sirupsen/logrus"
)

type SupportHandler struct {
	supportService *services.SupportService
}

func NewSupportHandler(supportService *services.SupportService) *SupportHandler {
	return &SupportHandler{
		supportService: supportService,
	}
}

func (h *SupportHandler) NewSupport(ctx context.Context, req *pb.NewSupportRequest) (*pb.NewSupportResponse, error) {

	supportID, err := h.supportService.CreateSupport(ctx, req.Name)
	if err != nil {
		return &pb.NewSupportResponse{
			Success:   false,
			Error:     err.Error(),
			SupportId: "",
		}, err
	}

	return &pb.NewSupportResponse{
		Success:   true,
		Error:     "",
		SupportId: supportID.String(),
	}, nil

}

func (h *SupportHandler) GetSupportTickets(ctx context.Context, req *pb.GetSupportTicketsRequest) (*pb.GetSupportTicketsResponse, error) {

	tickets, err := h.supportService.GetSupportTickets(ctx, req.SupportId)
	if err != nil {
		return &pb.GetSupportTicketsResponse{
			Success: false,
			Error:   err.Error(),
			Tickets: transformers.TicketModelToGRPC(tickets),
		}, err
	}

	logrus.Infof("grpc : %v", tickets)
	logrus.Infof("model : %v", transformers.TicketModelToGRPC(tickets))

	return &pb.GetSupportTicketsResponse{
		Success: true,
		Error:   "",
		Tickets: transformers.TicketModelToGRPC(tickets),
	}, nil

}
