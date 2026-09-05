package models

type TicketStatus int

const (
	Open TicketStatus = iota
	Answered
	Closed
)

type Ticket struct {
	ID        string
	UserID    string
	SupportID string
	Request   string
	Responses []string
	Status    TicketStatus
}
