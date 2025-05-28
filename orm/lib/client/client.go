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
	db      *sql.DB
	config  config.ORMConfig
	tables  Tables
	adapter adapter.DbAdapter
}

func CreateClient(config config.ORMConfig, currentAdapter adapter.DbAdapter) DbClient {
	connStr := currentAdapter.GetConnString(config)
	db, err := sql.Open(currentAdapter.GetDbDriver(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	tableMap := make(Tables)

	return DbClient{
		db:      db,
		config:  config,
		tables:  tableMap,
		adapter: currentAdapter,
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

func (dc DbClient) GetAdapter() adapter.DbAdapter {
	return dc.adapter
}

func (dc DbClient) Query(query string, args ...any) (*sql.Rows, error) {
	return dc.db.Query(query, args)
}
