package models

import "uuid"

type TicketStatus int

const (
	Opened TicketStatus = iota
	Answered
	Closed
)

type Ticket struct {
	ID        uuid.UUID    `db:"id"`
	UserID    uuid.UUID    `db:"user_id"`
	SupportID uuid.UUID    `db:"support_id"`
	Body      string       `db:"body"`
	Status    TicketStatus `db:"status"`
}

func TicketStatusToString(status TicketStatus) string {
	switch status {
	case Opened:
		return "Opened"
	case Answered:
		return "Answered"
	case Closed:
		return "Closed"
	}

	return ""
}
