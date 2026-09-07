package models

import (
	"time"
)

type Support struct {
	ID                      string    `db:"id"`
	CurrentAssignedTicketID string    `db:"current_assigned_ticket_id"`
	LastAssignedTicketTime  time.Time `db:"last_assigned_ticket_time"`
}
