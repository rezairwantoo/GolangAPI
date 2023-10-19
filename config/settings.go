package config

import (
	"case2/model"
	"case2/repository"
	"context"

	"github.com/pkg/errors"
)

var driverName string

type ISettings interface {
	Setup(s *Settings)
}
type Settings struct {
	ctx                 context.Context
	Config              *model.Config
	PostgresSQLProvider *repository.PostgresRepository
}

// NewSettings ...
func NewSettings(ctx context.Context) (*Settings, error) {
	conf := &Config{}
	if err := conf.NewConfig(ctx); err != nil {
		return nil, errors.New(err.Error())
	}

	return &Settings{
		ctx:    ctx,
		Config: conf.Config,
	}, nil
}

// Setup ...
func (*Settings) Setup(_ *Settings) {}

// Load ...
func (r *Settings) Load(x ISettings) {
	x.Setup(r)
}

func (*Settings) SetPostgresRepo(s ISettings) ISettings {
	return &Postgres{Setting: s}
}
