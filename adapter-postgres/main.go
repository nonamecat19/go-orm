package adapter_postgres

import (
	base "adapter-base"
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

func (ap AdapterPostgres) DeleteFromTable(tableName string) string {
	return base.DeleteFromTable(tableName)
}

func (ap AdapterPostgres) NormalizeSqlWithArgs(sql string, args []interface{}) string {
	return base.NormalizeSqlWithArgs(sql, args)
}
