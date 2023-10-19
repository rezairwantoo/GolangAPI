package config

import (
	"case2/model"
	"case2/repository"
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Setting ISettings
}

func (r *Postgres) Setup(s *Settings) {
	var dbConn *sqlx.DB
	dbConn = NewPostgresConnection(s.ctx, &s.Config.Postgres)
	if s.PostgresSQLProvider == nil {
		repositoryPostgres := repository.NewRepository(dbConn)
		s.PostgresSQLProvider = repositoryPostgres
	} else {
		s.PostgresSQLProvider.Conn = dbConn
	}

	r.Setting.Setup(s)
}

// NewPostgresConnection ...
func NewPostgresConnection(ctx context.Context, s *model.PostgreSQLConfig) *sqlx.DB {
	var err error
	postgresDriverName := "postgres"
	connection := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", s.Username, s.Password, s.Host, s.Port, s.Database)
	val := url.Values{}
	val.Add("TimeZone", "Asia/Jakarta")
	val.Add("sslmode", "disable")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	sqldb, err := sql.Open(postgresDriverName, dsn)
	if err != nil {
		err = errors.Wrap(err, "failed connect postgres")
		panic(err)
	}

	db := sqlx.NewDb(sqldb, postgresDriverName)
	if err = db.Ping(); err != nil {
		panic(errors.Wrap(err, "ping database postgres"))
	}

	db.SetMaxOpenConns(s.MaxOpenConns)                                    // The default is 0 (unlimited)
	db.SetMaxIdleConns(s.MaxIdleConns)                                    // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(s.ConnMaxLifetime) * time.Minute) // 0, connections are reused forever.

	return db
}
