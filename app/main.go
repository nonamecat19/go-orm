package main

import (
	adaptermssql "adapter-mssql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/core/utils"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	querybuilder "github.com/nonamecat19/go-orm/orm/lib/querybuilder"
)

func main() {
	//postgresConfig := config.ORMConfig{
	//	Host:     "127.0.0.1",
	//	Port:     15432,
	//	User:     "postgres",
	//	Password: "root",
	//	DbName:   "orm",
	//	SSLMode:  false,
	//}
	//
	//postgresAdapter := adapterpostgres.AdapterPostgres{}
	//
	//client := client2.CreateClient(postgresConfig, postgresAdapter)

	//sqliteConfig := config.ORMConfig{
	//	Path: "./sqlite.sqlite",
	//}
	//
	//sqliteAdapter := adaptersqlite.AdapterSQLite{}

	//client := client2.CreateClient(sqliteConfig, sqliteAdapter)

	mssqlConfig := config.ORMConfig{
		Host:     "127.0.0.1",
		Port:     1433,
		User:     "sa",
		Password: "1StrongPwd!!",
		DbName:   "master",
		SSLMode:  false,
	}

	mssqlAdapter := adaptermssql.AdapterMSSQL{}

	client := client2.CreateClient(mssqlConfig, mssqlAdapter)

	//mysqlConfig := config.ORMConfig{
	//	Host:     "127.0.0.1",
	//	Port:     3306,
	//	User:     "admin",
	//	Password: "root",
	//	DbName:   "orm",
	//}
	//
	//mysqlAdapter := adaptermysql.AdapterMySQL{}
	//
	//client := client2.CreateClient(mysqlConfig, mysqlAdapter)

	//users := []entities.User{
	//	{
	//		Name:   "test",
	//		Email:  "email@gmail.com",
	//		Gender: "male",
	//	},
	//	{
	//		Name:   "test2",
	//		Email:  "email2@gmail.com",
	//		Gender: "female",
	//	},
	//}
	//
	//err := querybuilder.CreateQueryBuilder(client).
	//	Debug().
	//	InsertMany(users)

	//err := querybuilder.CreateQueryBuilder(client).
	//	Debug().
	//	InsertOne(users[0])

	var users []entities.User

	err := querybuilder.CreateQueryBuilder(client).
		Where("name <> ? OR name <> ?", "test1", "User 200").
		AndWhere("name <> '2'").
		AndWhere("name <> ?", '3').
		Preload("orders").
		Preload("role").
		OrderBy("id DESC").
		Limit(8).
		Offset(2).
		FindMany(&users)

	//err := querybuilder.CreateQueryBuilder(client).
	//	Where("id = ?", 35).
	//	DeleteMany(&entities.User{})
	//
	//err := querybuilder.CreateQueryBuilder(client).
	//	Debug().
	//	SetValues(map[string]any{"name": "test"}).
	//	Where("id > ?", 32).
	//	AndWhere("id < 42").
	//	UpdateMany(&entities.User{})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.PrintStructSlice(users)

	//var orders []entities.Order
	//
	//err := querybuilder.CreateQueryBuilder(client).
	//	Where("id <> ?", 8).
	//	AndWhere("count <> ?", 7).
	//	Debug().
	//	OrderBy("id ASC").
	//	Preload("user").
	//	Limit(15).
	//	Offset(1).
	//	FindMany(&orders)
	//
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//utils.PrintStructSlice(orders)
}
