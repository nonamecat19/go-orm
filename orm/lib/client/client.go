package client

import (
	"database/sql"
	"github.com/nonamecat19/go-orm/core/lib/adapter"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/core/lib/scheme"
	"log"
)

type Tables = map[string]scheme.TableScheme

type DbClient struct {
	db     *sql.DB
	config config.ORMConfig
	tables Tables
}

func CreateClient(config config.ORMConfig, currentAdapter adapter.Adapter) DbClient {
	connStr := currentAdapter.GetConnString(config)
	db, err := sql.Open(currentAdapter.GetDbDriver(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	tableMap := make(Tables)

	return DbClient{
		db:     db,
		config: config,
		tables: tableMap,
	}
}

func (dc DbClient) GetDb() *sql.DB {
	return dc.db
}

func (dc DbClient) GetConfig() config.ORMConfig {
	return dc.config
}

func (dc DbClient) GetTables() Tables {
	return dc.tables
}

func (dc DbClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return dc.db.Query(query, args)
}
