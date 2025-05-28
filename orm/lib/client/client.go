package client

import (
	"database/sql"
	"github.com/nonamecat19/go-orm/core/lib/adapter"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"log"
)

type DbClient struct {
	db      *sql.DB
	config  config.ORMConfig
	adapter adapter.DbAdapter
}

func CreateClient(config config.ORMConfig, currentAdapter adapter.DbAdapter) DbClient {
	connStr := currentAdapter.GetConnString(config)
	db, err := sql.Open(currentAdapter.GetDbDriver(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	return DbClient{
		db:      db,
		config:  config,
		adapter: currentAdapter,
	}
}

func (dc DbClient) GetDb() *sql.DB {
	return dc.db
}

func (dc DbClient) GetConfig() config.ORMConfig {
	return dc.config
}

func (dc DbClient) GetAdapter() adapter.DbAdapter {
	return dc.adapter
}

func (dc DbClient) Query(query string, args ...any) (*sql.Rows, error) {
	return dc.db.Query(query, args)
}
