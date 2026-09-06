package models

type User struct {
	ID      string   `json:"id"`
	Tickets []Ticket `json:"tickets"`
}
