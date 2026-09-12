package repositories

import (
	"context"
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

func (r *TicketRepo) NewTicket(ctx context.Context, userID, body string) (*models.Ticket, error) {
	var ticket models.Ticket
	err := r.dbConn.GetContext(ctx, &ticket, "INSERT INTO tickets(user_id,body) VALUES ($1,$2) RETURNING *", userID, body)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (r *TicketRepo) GetTicketsWithStatus(ctx context.Context, status string) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	err := r.dbConn.SelectContext(ctx, &tickets, "SELECT * FROM tickets WHERE status=$1", status)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepo) CloseTicket(ctx context.Context, ticketID string, supportID string) error {

	tx, err := r.dbConn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = r.dbConn.ExecContext(ctx, "UPDATE tickets SET status = 'Closed' WHERE id=$1 AND support_id=$2;", ticketID, supportID)
	if err != nil {
		return err
	}

	_, err = r.dbConn.ExecContext(ctx, "UPDATE supports SET current_ticket_id = NULL WHERE support_id=$1;", supportID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *TicketRepo) TransferTicket(ctx context.Context, ticketID, newSupportID string) error {
	_, err := r.dbConn.ExecContext(ctx, "UPDATE tickets SET support_id=$1 WHERE id=$2;", newSupportID, ticketID)
	return err
}

func (r *TicketRepo) AnswerTicket(ctx context.Context, ticketID, supportID, body string) error {

	tx, err := r.dbConn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.QueryContext(ctx, "INSERT INTO ticket_responses(ticket_id,body) VALUES ($1,$2)", ticketID, body)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE tickets SET status=$1 WHERE ticket_id=$2;", models.Answered, ticketID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
