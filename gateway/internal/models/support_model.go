package models

import (
	"time"
)

type Support struct {
	ID                      string
	CurrentAssignedTicketID string
	lastAssignedTicketTime  time.Time
}
