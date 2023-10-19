package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	ErrFailedQuery        = "failed query"
	ErrFailedPrepareQuery = "failed prepare query"
)

type Querier interface {
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
}

// newQueryable choose transaction over database connection if found one
func newQueryable(db *sqlx.DB, tx *sqlx.Tx) Querier {
	if tx != nil {
		return tx
	}
	return db
}

// ExecStatementContext is a helper function to execute insert, update, delete query
func ExecStatementContext(ctx context.Context, conn Querier, rawQuery string, args map[string]any) (sql.Result, error) {
	stmt, err := conn.PrepareNamedContext(ctx, rawQuery)
	if err != nil {
		return nil, errors.Wrap(err, ErrFailedPrepareQuery)
	}
	defer func() {
		err = stmt.Close()
	}()

	res, err := stmt.ExecContext(ctx, args)
	if err != nil {
		return nil, errors.Wrap(err, ErrFailedQuery)
	}

	return res, nil
}

func GetStatementContext[T any](ctx context.Context, conn Querier, rawQuery string, args map[string]any) (T, error) {
	var destination T
	stmt, err := conn.PrepareNamedContext(ctx, rawQuery)
	if err != nil {
		return destination, errors.Wrap(err, ErrFailedPrepareQuery)
	}
	defer func() {
		err = stmt.Close()
	}()
	if err = stmt.GetContext(ctx, &destination, args); err != nil {
		return destination, errors.Wrap(err, ErrFailedQuery)
	}
	return destination, nil
}

// SelectStatementContext is a helper function to execute select query e.g. if you want to return array value
func SelectStatementContext[T any](ctx context.Context, conn Querier, rawQuery string, args map[string]any) ([]T, error) {
	var destination []T
	stmt, err := conn.PrepareNamedContext(ctx, rawQuery)
	if err != nil {
		return destination, errors.Wrap(err, ErrFailedPrepareQuery)
	}
	defer func() {
		err = stmt.Close()
	}()
	if err = stmt.SelectContext(ctx, &destination, args); err != nil {
		return destination, errors.Wrap(err, ErrFailedQuery)
	}
	return destination, nil
}
