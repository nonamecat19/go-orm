package main

import (
	"github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/orm/lib/migrate"
	"log"
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

	appEntities := []coreEntities.IEntity{
		entities.Task{},
		entities.User{},
	}

	migrate.PushEntity(ormConfig, appEntities)

	client := client2.CreateClient(ormConfig)
	data, err := client.Read(&entities.User{})
	if err != nil {
		panic(err)
	}
	log.Println(data)

}
