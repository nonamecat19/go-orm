package tests

import (
	adaptermysql "adapter-mysql"
	adapterpostgres "adapter-postgres"
	"github.com/nonamecat19/go-orm/core/lib/config"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
)

func GetPostgresTestClient() client2.DbClient {
	postgresConfig := config.ORMConfig{
		Host:     "127.0.0.1",
		Port:     15432,
		User:     "postgres",
		Password: "root",
		DbName:   "orm",
		SSLMode:  false,
	}

	postgresAdapter := adapterpostgres.AdapterPostgres{}
	return client2.CreateClient(postgresConfig, postgresAdapter)
}

func GetMySQLTestClient() client2.DbClient {
	mysqlConfig := config.ORMConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "admin",
		Password: "root",
		DbName:   "orm",
	}

	mysqlAdapter := adaptermysql.AdapterMySQL{}

	return client2.CreateClient(mysqlConfig, mysqlAdapter)
}
