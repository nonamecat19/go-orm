package main

import (
	adapterpostgres "adapter-postgres"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/core/utils"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	querybuilder "github.com/nonamecat19/go-orm/orm/lib/querybuilder"
)

func main() {
	ormConfig := config.ORMConfig{
		Host:     "127.0.0.1",
		Port:     15432,
		User:     "postgres",
		Password: "root",
		DbName:   "orm",
		SSLMode:  false,
	}

	adapter := adapterpostgres.AdapterPostgres{}

	client := client2.CreateClient(ormConfig, adapter)

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

	//var users []entities.User

	//err := querybuilder.CreateQueryBuilder(client).
	//	//Where("name <> ? OR name <> ?", "test1", "User 200").
	//	//AndWhere("name <> '2'").
	//	//AndWhere("name <> ?", '3').
	//	//Select("users.id", "users.name", "users.created_at").
	//	Debug().
	//	Preload("orders").
	//	//AndWhere("name <> ?", "User 200").
	//	//OrderBy("id DESC").
	//	//Limit(3).
	//	//LeftJoinAndSelect("orders", "users.id = orders.user_id", "orders.id", "orders.count").
	//	//Offset(10).
	//	FindMany(&users)

	//err := querybuilder.CreateQueryBuilder(client).
	//	Where("id = ?", 35).
	//	DeleteMany(&entities.User{})
	//
	//err := querybuilder.CreateQueryBuilder(client).
	//	Debug().
	//	SetValues(map[string]interface{}{"name": "test"}).
	//	Where("id > ?", 32).
	//	AndWhere("id < 42").
	//	UpdateMany(&entities.User{})

	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//utils.PrintStructSlice(users)

	var orders []entities.Order

	err := querybuilder.CreateQueryBuilder(client).
		Where("id <> ?", 8).
		AndWhere("count <> ?", 7).
		Debug().
		OrderBy("id ASC").
		Preload("user").
		Limit(15).
		Offset(1).
		FindMany(&orders)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.PrintStructSlice(orders)
}
