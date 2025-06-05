package adapter_mysql

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/lib/config"
)

type AdapterMySQL struct{}

func (ap AdapterMySQL) GetConnString(config config.ORMConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.DbName)
}

func (ap AdapterMySQL) GetDbDriver() string {
	return "mysql"
}
