package models

import (
	"time"
	"uuid"
)

type TicketResponse struct {
	CreatedAt time.Time `db:"created_at"`
	TicketID  uuid.UUID    `db:"ticket_id"`
	body      string    `db:"body"`
}
