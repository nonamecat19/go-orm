package adapter_sqlite

import (
	"github.com/nonamecat19/go-orm/core/lib/config"
)

type AdapterSQLite struct{}

func (ap AdapterSQLite) GetConnString(config config.ORMConfig) string {
	return config.Path
}

func (ap AdapterSQLite) GetDbDriver() string {
	return "sqlite3"
}
