package repositories

import (
	"context"
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

func (r *UserRepo) CreateUser(ctx context.Context, name string) (*models.User, error) {
	var user models.User
	err := r.dbConn.GetContext(ctx, &user, "INSERT INTO users (name) VALUES ($1) RETURNING *", name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserTickets(ctx context.Context, userID string) ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.dbConn.SelectContext(ctx, &tickets, "SELECT * FROM tickets WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}
