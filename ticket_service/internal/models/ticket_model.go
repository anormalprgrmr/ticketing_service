package models

type TicketStatus int

const (
	Open TicketStatus = iota
	Answered
	Closed
)

type Ticket struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	SupportID string       `json:"support_id"`
	Body      string       `json:"body"`
	Responses []string     `json:"responses"`
	Status    TicketStatus `json:"status"`
}
