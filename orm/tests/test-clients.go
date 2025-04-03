package tests

import (
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
