package repositories

import (
	"ticket_service/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	dbConn *sqlx.DB
}

func NewUserRepo(dbConn *sqlx.DB) *UserRepo {
	return &UserRepo{
		dbConn: dbConn,
	}
}

func (r *UserRepo) CreateUser(name string) (*models.User, error) {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil, nil
}

func (r *UserRepo) GetMyTickets(userID string) ([]models.Ticket, error) {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil, nil
}
