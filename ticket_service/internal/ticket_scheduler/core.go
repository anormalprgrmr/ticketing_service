package ticketscheduler

import (
	"ticket_service/internal/models"
)

type TicketScheduler struct {
	TicketQueue chan models.Ticket
}

func (ts *TicketScheduler) Trigger(ticket *models.Ticket) {

}
