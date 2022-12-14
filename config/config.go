package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port              int    `env:"PORT"`
	MySQLUser         string `env:"MYSQL_USER"`
	MYSQLRootPassword string `env:"MYSQL_ROOT_PASSWORD"`
	MYSQLAddr         string `env:"MYSQL_ADDR"`
	MYSQLDbName       string `env:"MYSQL_DATABASE"`
	RedisHost         string `env:"REDIS_HOST"`
	RedisPort         int    `env:"REDIS_PORT"`
	Auth0Domain       string `env:"AUTH0_DOMAIN"`
	Auth0Audience     string `env:"AUTH0_AUDIENCE"`
	FrontendHost      string `env:"FRONTEND_HOST"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
