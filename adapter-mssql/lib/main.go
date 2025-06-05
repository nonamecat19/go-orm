package adapter_mssql

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/config"
)

type AdapterMSSQL struct{}

func (ap AdapterMSSQL) GetConnString(config config.ORMConfig) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		config.User, config.Password, config.Host, config.Port, config.DbName)
}

func (ap AdapterMSSQL) GetDbDriver() string {
	return "sqlserver"
}
