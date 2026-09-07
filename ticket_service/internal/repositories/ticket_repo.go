package repositories

import (
	"ticket_service/internal/models"

	"github.com/jmoiron/sqlx"
)

type TicketRepo struct {
	dbConn *sqlx.DB
}

func NewTicketRepo(dbConn *sqlx.DB) *TicketRepo {
	return &TicketRepo{
		dbConn: dbConn,
	}
}

func (r *TicketRepo) NewTicket(userID, body string) error {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil
}

func (r *TicketRepo) GetTicketByID(userID, ticketID string) (*models.Ticket, error) {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil, nil
}

func (r *TicketRepo) ChangeTicketStatus(ticketID string, status models.TicketStatus) error {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil
}

func (r *TicketRepo) ResponseTicket(ticketID, body string) error {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil
}

func (r *TicketRepo) GetTicketsWithStatus(status models.TicketStatus) ([]models.Ticket, error) {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil, nil
}

func (r *TicketRepo) TransferTicket(ticketID, newSupportID string) error {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil
}
