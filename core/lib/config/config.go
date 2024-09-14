package config

import "github.com/nonamecat19/go-orm/core/lib/scheme"

type ORMConfig struct {
	DbDriver string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SSLMode  bool
	Tables   []scheme.TableScheme
}
