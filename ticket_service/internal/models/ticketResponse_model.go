package models

import "time"

type TicketResponse struct {
	CreatedAt time.Time `db:"created_at"`
	TicketID  string    `db:"ticket_id"`
	body      string    `db:"body"`
}
