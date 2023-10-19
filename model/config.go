package model

// Config ...
type Config struct {
	App      AppConfig        `json:"app" validate:"required"`
	Postgres PostgreSQLConfig `json:"postgres" validate:"required"`
}

// AppConfig ...
type AppConfig struct {
	Port string `json:"port" validate:"required"`
}

type PostgreSQLConfig struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	MaxOpenConns    int    `json:"max_open_conns"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
	Database        string `json:"dbname"`
}
