package config

import (
	"case2/model"
	"case2/model/constant"
	"context"

	"github.com/pkg/errors"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Config *model.Config
}

func LoadConfigFile(ctx context.Context) {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		zlog.Info().Interface("error", err).Msg("Can't find config file")
	}
}

func (c *Config) NewConfig(ctx context.Context) (err error) {
	c.Config = &model.Config{}
	if err = c.setConfigApp(); err != nil {
		return errors.New("failed set config app")
	}
	c.SetConfigPostgres(ctx)
	return
}

func (c *Config) setConfigApp() error {
	appGrpcPort := viper.GetString(constant.AppGrpcPort)

	c.Config.App = model.AppConfig{
		Port: appGrpcPort,
	}

	return nil
}

func (c *Config) SetConfigPostgres(ctx context.Context) error {
	var err error
	if err = viper.UnmarshalKey("db-postgres", &c.Config.Postgres); err != nil {
		return errors.Wrap(err, "failed init config")
	}
	return nil
}
