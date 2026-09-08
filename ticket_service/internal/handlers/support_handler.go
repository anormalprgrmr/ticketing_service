package handlers

import (
	"context"
	pb "ticket_service/internal/protos"
	"ticket_service/internal/services"
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
