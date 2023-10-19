package model

import "time"

type Products struct {
	ID        int64      `db:"id"`
	Item      string     `db:"item"`
	Price     int64      `db:"price"`
	IsActive  bool       `db:"is_active"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type ProductsTotalData struct {
	Total int64 `db:"total"`
}
