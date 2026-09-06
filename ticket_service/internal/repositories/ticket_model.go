package repositories

import "database/sql"

type TicketRepo struct {
	dbConn *sql.DB
}

func NewTicketRepo(dbConn *sql.DB) *TicketRepo {
	return &TicketRepo{}
}

func (r *TicketRepo) NewTicket(userID, body string) error {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
}
