package ticketscheduler

import (
	"context"
	"database/sql"
	"errors"
	"sync"
	"ticket_service/internal/event"
	"ticket_service/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TicketScheduler struct {
	mutex       sync.Mutex
	ticketQueue Queue[models.Ticket]
	dbConn      *sqlx.DB
	eb          event.EventSignalBus
}

func NewTicketScheduler(ctx context.Context, dbConn *sqlx.DB, eb event.EventSignalBus) (*TicketScheduler, error) {
	ts := &TicketScheduler{
		mutex:       sync.Mutex{},
		ticketQueue: Queue[models.Ticket]{},
		dbConn:      dbConn,
		eb:          eb,
	}

	err := ts.start(ctx)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (ts *TicketScheduler) start(ctx context.Context) error {
	// Recover all existing/unprocessed work first.
	if err := ts.syncAll(ctx); err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case <-ts.eb.Subscribe():
				log.Debug("🦓🦓 event triggered the scheduler")

				if _, err := ts.updateScheduler(ctx); err != nil {
					log.Errorf("error updating scheduler: %v", err)
				}
			}
		}
	}()

	return nil
}

func (ts *TicketScheduler) updateScheduler(ctx context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx, err := ts.dbConn.BeginTxx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	var support models.Support
	err = tx.GetContext(
		ctx,
		&support,
		`SELECT *
         FROM supports
         WHERE current_ticket_id IS NULL
         ORDER BY last_assigned_ticket_time
         LIMIT 1`,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	var ticket models.Ticket
	err = tx.GetContext(
		ctx,
		&ticket,
		`SELECT *
         FROM tickets
         WHERE support_id IS NULL
           AND status = 'Opened'
         ORDER BY created_at
         LIMIT 1`,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	_, err = tx.ExecContext(
		ctx,
		`UPDATE tickets SET support_id = $1 WHERE id = $2`,
		support.ID,
		ticket.ID,
	)
	if err != nil {
		return false, err
	}

	_, err = tx.ExecContext(
		ctx,
		`UPDATE supports SET current_ticket_id = $1 WHERE id = $2`,
		ticket.ID,
		support.ID,
	)
	if err != nil {
		return false, err
	}

	if err := tx.Commit(); err != nil {
		return false, err
	}

	return true, nil
}

func (ts *TicketScheduler) syncAll(ctx context.Context) error {
	for {
		updated, err := ts.updateScheduler(ctx)
		if err != nil {
			return err
		}

		if !updated {
			return nil
		}
	}
}
