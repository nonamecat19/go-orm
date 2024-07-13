package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"orm/scheme"
)

type DbClient struct {
	db     *sql.DB
	config ORMConfig
	tables map[string]scheme.TableScheme
}

func (c *DbClient) table(table string) scheme.TableScheme {
	return c.tables[table]
}

type ORMConfig struct {
	DbDriver string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SSLMode  bool
	Tables   []scheme.TableScheme
}

func createClient(config ORMConfig) DbClient {
	var connStr string
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
	db, err := sql.Open(config.DbDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db)

	tableMap := make(map[string]scheme.TableScheme)
	for _, table := range config.Tables {
		tableMap[table.Name] = table
	}

	return DbClient{
		db:     db,
		config: config,
		tables: tableMap,
	}
}

func main() {
	userTableSchema := scheme.TableScheme{
		Name: "users",
		Fields: []scheme.Field{
			{Name: "id", Type: "integer", Nullability: false},
			{Name: "name", Type: "varchar(50)", Nullability: true},
			{Name: "email", Type: "varchar(100)", Nullability: false},
			{Name: "password", Type: "varchar(100)", Nullability: true},
		},
	}

	config := ORMConfig{
		DbDriver: "postgres",
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "root",
		DbName:   "orm",
		SSLMode:  false,
		Tables: []scheme.TableScheme{
			userTableSchema,
		},
	}

	db := createClient(config)

	db.table("users").Create(db.db, "name", "email", "password")

	//err := schema.createTable(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result, err := schema.create(db, "John Doe", "john@example.com", "password123")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result)
	//
	//users, err := schema.read(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, user := range users {
	//	fmt.Println(user)
	//}
	//
	//result, err = schema.update(db, 1, "Jane Doe", "jane@example.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result)
	//
	//result, err = schema.delete(db, 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result)
}
