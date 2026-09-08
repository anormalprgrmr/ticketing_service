package models

import (
	"time"
	"uuid"
)

type Support struct {
	ID                     uuid.UUID `db:"id"`
	Name                   string    `db:"name"`
	LastAssignedTicketTime time.Time `db:"last_assigned_ticket_time"`
}
