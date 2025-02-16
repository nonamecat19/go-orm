package client

import (
	"database/sql"
	"fmt"
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

func CreateClient(config config.ORMConfig) DbClient {
	var connStr string
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
	db, err := sql.Open(config.DbDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//}(db)

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

//func (dc DbClient) Read() ([]entities2.User, error) {
//	query := "SELECT name FROM users;"
//	rows, err := dc.db.Query(query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var users []entities2.User
//	for rows.Next() {
//		var user entities2.User
//		err = rows.Scan(&user.Name)
//		if err != nil {
//			return nil, err
//		}
//		users = append(users, user)
//	}
//
//	return users, nil
//}
