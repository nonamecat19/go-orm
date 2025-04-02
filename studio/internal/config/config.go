package config

import (
	"github.com/nonamecat19/go-orm/core/lib/config"
)

type Config struct {
	ServerAddr string
}

func NewConfig() Config {
	return Config{
		ServerAddr: ":8080",
	}
}

var PostgresConfig = config.ORMConfig{
	Host:     "localhost",
	Port:     15432,
	User:     "postgres",
	Password: "root",
	DbName:   "orm",
	SSLMode:  false,
}
