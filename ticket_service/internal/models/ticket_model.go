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
	Status    TicketStatus `db:"status"`
}
