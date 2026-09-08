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
			Status:    resolveTicketStatus(ticket.Status),
		})
	}

	return pbTickets
}

func resolveTicketStatus(status models.TicketStatus) pb.TicketStatus {
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
