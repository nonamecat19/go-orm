package test_utils

import (
	adaptermssql "adapter-mssql"
	adaptermysql "adapter-mysql"
	adapterpostgres "adapter-postgres"
	adaptersqlite "adapter-sqlite"
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

func GetSQLiteTestClient() client2.DbClient {
	sqliteConfig := config.ORMConfig{
		Path: "../../../sqlite.sqlite",
	}

	sqliteAdapter := adaptersqlite.AdapterSQLite{}

	return client2.CreateClient(sqliteConfig, sqliteAdapter)
}

func GetMSSQLTestClient() client2.DbClient {
	mssqlConfig := config.ORMConfig{
		Host:     "127.0.0.1",
		Port:     1433,
		User:     "sa",
		Password: "1StrongPwd!!",
		DbName:   "master",
		SSLMode:  false,
	}

	mssqlAdapter := adaptermssql.AdapterMSSQL{}

	return client2.CreateClient(mssqlConfig, mssqlAdapter)
}
