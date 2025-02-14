package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	querybuilder "github.com/nonamecat19/go-orm/orm/lib/querybuilder"
)

func main() {
	ormConfig := config.ORMConfig{
		DbDriver: "postgres",
		Host:     "127.0.0.1",
		Port:     15432,
		User:     "postgres",
		Password: "root",
		DbName:   "orm",
		SSLMode:  false,
	}

	//appEntities := []coreEntities.IEntity{
	//	entities.Task{},
	//	entities.User{},
	//}

	//migrate.PushEntity(ormConfig, appEntities)

	client := client2.CreateClient(ormConfig)
	qb := querybuilder.CreateQueryBuilder(client)

	user := &entities.User{}
	query, err := qb.
		Where("name = ?", "John Doe").
		OrderBy("id DESC").
		Limit(10).
		FindMany(user)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Built Query:", *query)
}
