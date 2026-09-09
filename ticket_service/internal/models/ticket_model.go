package models

import (
	"time"
	"uuid"
)

type TicketStatus string

const (
	Opened   TicketStatus = "Opened"
	Answered TicketStatus = "Answered"
	Closed   TicketStatus = "Closed"
)

type Ticket struct {
	ID        uuid.UUID    `db:"id"`
	UserID    uuid.UUID    `db:"user_id"`
	SupportID *uuid.UUID   `db:"support_id"`
	Body      string       `db:"body"`
	Status    TicketStatus `db:"status"`
	CreatedAt time.Time    `db:"created_at"`
}
