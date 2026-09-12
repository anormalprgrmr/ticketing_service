package transformers

import (
	"ticket_service/internal/models"
	pb "ticket_service/internal/protos"
)

func TicketModelToGRPC(tickets []models.Ticket) []*pb.Ticket {

	pbTickets := make([]*pb.Ticket, len(tickets))

	for _, ticket := range tickets {
		pbTickets = append(pbTickets, &pb.Ticket{
			Id:        ticket.ID.String(),
			UserId:    ticket.UserID.String(),
			SupportId: ticket.SupportID.String(),
			Body:      ticket.Body,
			Status:    ConvertTicketStatusModelToGRPC(ticket.Status),
		})
	}

	return pbTickets
}

func ConvertTicketStatusModelToGRPC(status models.TicketStatus) pb.TicketStatus {
	switch status {
	case models.Opened:
		return pb.TicketStatus_TICKET_STATUS_OPEN
	case models.Answered:
		return pb.TicketStatus_TICKET_STATUS_ANSWERED
	case models.Closed:
		return pb.TicketStatus_TICKET_STATUS_CLOSED
	default:
		return pb.TicketStatus_TICKET_STATUS_UNSPECIFIED
	}
}

func ConvertTicketStatusGRPCToModel(status pb.TicketStatus) models.TicketStatus {
	switch status {
	case pb.TicketStatus_TICKET_STATUS_OPEN:
		return models.Opened
	case pb.TicketStatus_TICKET_STATUS_ANSWERED:
		return models.Answered
	case pb.TicketStatus_TICKET_STATUS_CLOSED:
		return models.Closed
	}

	return models.Opened
}
