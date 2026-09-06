package models

import (
	"time"
)

type Support struct {
	ID                      string    `json:"id"`
	CurrentAssignedTicketID string    `json:"current_assigned_ticket_id"`
	LastAssignedTicketTime  time.Time `json:"last_assigned_ticket_time"`
}
