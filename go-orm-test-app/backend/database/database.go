package database

import (
	_ "github.com/lib/pq"
	adapter "github.com/nonamecat19/go-orm/adapter-postgres/lib"
	"github.com/nonamecat19/go-orm/core/lib/config"
	client "github.com/nonamecat19/go-orm/orm/lib/client"
	"os"
	"strconv"
)

var DbClient client.DbClient

func InitDbClient() {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dbConfig := config.ORMConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  false,
	}

	DbClient = client.CreateClient(dbConfig, adapter.AdapterPostgres{})
}
