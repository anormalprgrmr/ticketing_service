package ticketscheduler

import (
	"context"
	"sync"
	"ticket_service/internal/models"

	"github.com/jmoiron/sqlx"
)

type TicketScheduler struct {
	mutex       sync.Mutex
	ticketQueue Queue[models.Ticket]
	dbConn      *sqlx.DB
}

func NewTicketScheduler(dbConn *sqlx.DB) *TicketScheduler {
	return &TicketScheduler{
		mutex:       sync.Mutex{},
		ticketQueue: Queue[models.Ticket]{},
		dbConn:      dbConn,
	}
}

func (ts *TicketScheduler) Trigger(ctx context.Context) error {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	tx, err := ts.dbConn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var support models.Support
	err = tx.GetContext(ctx, &support, "SELECT * FROM supports WHERE current_ticket_id IS NULL AND status='Opened' ORDER BY last_assigned_ticket_time LIMIT 1 ")
	if err != nil {
		return err
	}

	var ticket models.Ticket
	err = tx.GetContext(ctx, &ticket, "SELECT * FROM tickets WHERE support_id IS NULL AND status=Opened ORDER BY created_at LIMIT 1 ")
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE tickets SET support_id=$1 where id=$2", support.ID, ticket.ID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE supports SET current_ticket_id=$1 where id=$2", ticket.ID, support.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
