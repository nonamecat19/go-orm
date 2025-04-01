package adapter_postgres

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/config"
)

type AdapterPostgres struct{}

func (ap AdapterPostgres) GetConnString(config config.ORMConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
}

func (ap AdapterPostgres) GetDbDriver() string {
	return "postgres"
}
