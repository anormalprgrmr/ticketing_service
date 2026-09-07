package repositories

import (
	"context"
	"ticket_service/internal/models"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type SupportRepo struct {
	dbConn *sqlx.DB
}

func NewSupportRepo(dbConn *sqlx.DB) *SupportRepo {
	return &SupportRepo{
		dbConn: dbConn,
	}
}

func (r *SupportRepo) CreateSupport(ctx context.Context, name string) (supportID string, err error) {
	var support models.Support

	err = r.dbConn.GetContext(ctx, &support, "INSERT INTO supports (name) VALUES ($1)", name)
	if err != nil {
		return "", err
	}

	log.Debugf("inserted in DB : %v", support)

	return support.ID, nil
}

func (r *SupportRepo) GetSupportTickets(supportID string) (supportTickets []models.Ticket, err error) {
	var tickets []models.Ticket

	err = r.dbConn.Select(&tickets, "SELECT * FROM tickets WHERE support_id = $1", supportID)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}
