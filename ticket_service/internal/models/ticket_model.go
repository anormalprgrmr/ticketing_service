package models

type TicketStatus int

const (
	Open TicketStatus = iota
	Answered
	Closed
)

type Ticket struct {
	ID        string       `db:"id"`
	UserID    string       `db:"user_id"`
	SupportID string       `db:"support_id"`
	Body      string       `db:"body"`
	Responses []string     `db:"responses"`
	Status    TicketStatus `db:"status"`
}
