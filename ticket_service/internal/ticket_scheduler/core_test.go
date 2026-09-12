package ticketscheduler

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	t.Cleanup(func() {
		db.Close()
	})

	return sqlxDB, mock
}

func newScheduler(db *sqlx.DB) *TicketScheduler {
	ts, _ := NewTicketScheduler(context.Background(), db, nil)

	return ts
}

func TestUpdateScheduler_AssignsTicketToSupport(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	supportID := uuid.New()
	ticketID := uuid.New()

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID.String()),
		)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(ticketID.String()),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE tickets SET support_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(supportID, ticketID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE supports SET current_ticket_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(ticketID, supportID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectCommit()

	updated, err := scheduler.updateScheduler(context.Background())

	require.NoError(t, err)
	require.True(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateScheduler_NoAvailableSupport(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectRollback()

	updated, err := scheduler.updateScheduler(context.Background())

	require.NoError(t, err)
	require.False(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateScheduler_NoAvailableTicket(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	supportID := uuid.New()

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID.String()),
		)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectRollback()

	updated, err := scheduler.updateScheduler(context.Background())

	require.NoError(t, err)
	require.False(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateScheduler_SupportQueryError(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	expectedErr := errors.New("database error")

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnError(expectedErr)

	mock.ExpectRollback()

	updated, err := scheduler.updateScheduler(context.Background())

	require.ErrorIs(t, err, expectedErr)
	require.False(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateScheduler_TicketQueryError(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	supportID := uuid.New()
	expectedErr := errors.New("ticket query failed")

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID.String()),
		)

	// Ticket query fails.
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnError(expectedErr)

	mock.ExpectRollback()

	updated, err := scheduler.updateScheduler(context.Background())

	require.ErrorIs(t, err, expectedErr)
	require.False(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateScheduler_TicketUpdateError(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	supportID := uuid.New()
	ticketID := uuid.New()

	expectedErr := errors.New("failed to update ticket")

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID.String()),
		)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(ticketID.String()),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE tickets SET support_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(supportID, ticketID).
		WillReturnError(expectedErr)

	mock.ExpectRollback()

	updated, err := scheduler.updateScheduler(context.Background())

	require.ErrorIs(t, err, expectedErr)
	require.False(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateScheduler_SupportUpdateError(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	supportID := uuid.New()
	ticketID := uuid.New()

	expectedErr := errors.New("failed to update support")

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID.String()),
		)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(ticketID.String()),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE tickets SET support_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(supportID, ticketID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE supports SET current_ticket_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(ticketID, supportID).
		WillReturnError(expectedErr)

	mock.ExpectRollback()

	updated, err := scheduler.updateScheduler(context.Background())

	require.ErrorIs(t, err, expectedErr)
	require.False(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateScheduler_CommitError(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	supportID := uuid.New()
	ticketID := uuid.New()

	expectedErr := errors.New("commit failed")

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID.String()),
		)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(ticketID.String()),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE tickets SET support_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(supportID, ticketID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE supports SET current_ticket_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(ticketID, supportID).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectCommit().
		WillReturnError(expectedErr)

	updated, err := scheduler.updateScheduler(context.Background())

	require.ErrorIs(t, err, expectedErr)
	require.False(t, updated)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSyncAll_NoPendingWork(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectRollback()

	err := scheduler.syncAll(context.Background())

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSyncAll_ProcessesAllPendingTickets(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	supportID1 := uuid.New()
	ticketID1 := uuid.New()

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID1.String()),
		)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(ticketID1.String()),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE tickets SET support_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(supportID1, ticketID1).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE supports SET current_ticket_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(ticketID1, supportID1).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectCommit()

	supportID2 := uuid.New()
	ticketID2 := uuid.New()

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(supportID2.String()),
		)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM tickets
		 WHERE support_id IS NULL
		   AND status = 'Opened'
		 ORDER BY created_at
		 LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
			}).AddRow(ticketID2.String()),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE tickets SET support_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(supportID2, ticketID2).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectExec(
		regexp.QuoteMeta(
			`UPDATE supports SET current_ticket_id = $1 WHERE id = $2`,
		),
	).
		WithArgs(ticketID2, supportID2).
		WillReturnResult(
			sqlmock.NewResult(1, 1),
		)

	mock.ExpectCommit()

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectRollback()

	err := scheduler.syncAll(context.Background())

	require.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSyncAll_ReturnsError(t *testing.T) {
	db, mock := setupTestDB(t)

	scheduler := newScheduler(db)

	expectedErr := errors.New("database unavailable")

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT *
		 FROM supports
		 WHERE current_ticket_id IS NULL
		 ORDER BY last_assigned_ticket_time
		 LIMIT 1`)).
		WillReturnError(expectedErr)

	mock.ExpectRollback()

	err := scheduler.syncAll(context.Background())

	require.ErrorIs(t, err, expectedErr)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestNewTicketScheduler(t *testing.T) {
	db, _ := setupTestDB(t)

	scheduler, err := NewTicketScheduler(
		context.Background(),
		db,
		nil,
	)

	require.Nil(t, err)
	require.NotNil(t, scheduler)
	require.Equal(t, db, scheduler.dbConn)
	require.Nil(t, scheduler.eb)
}
