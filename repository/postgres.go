package repository

import (
	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	Conn *sqlx.DB
}

func NewRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}
