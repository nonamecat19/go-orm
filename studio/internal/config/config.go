package config

import "github.com/nonamecat19/go-orm/orm/lib/scheme"

type Config struct {
	ServerAddr string
}

func NewConfig() Config {
	return Config{
		ServerAddr: ":8080",
	}
}

var PostgresConfig = ORMConfig{
	DbDriver: "postgres",
	Host:     "localhost",
	Port:     15432,
	User:     "postgres",
	Password: "root",
	DbName:   "orm",
	SSLMode:  false,
	Tables:   []scheme.TableScheme{},
}
