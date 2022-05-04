package apiserver

import sql_user "github.com/krizons/rest-api-golang/internal/sql"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	DB       *sql_user.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		DB:       sql_user.NewConfig(),
	}
}
