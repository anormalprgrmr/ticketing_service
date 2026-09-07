package repositories

import "github.com/jmoiron/sqlx"

type SupportRepo struct {
	dbConn *sqlx.DB
}

func NewSupportRepo(dbConn *sqlx.DB) *SupportRepo {
	return &SupportRepo{
		dbConn: dbConn,
	}
}

func (r *SupportRepo) CreateSupport(name string) error {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil
}

func (r *SupportRepo) GetSupportTickets(name string) error {
	r.dbConn.Query("INSERT INTO tickets VALUES ($1)")
	return nil
}
