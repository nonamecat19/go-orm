package main

import (
	"fmt"
	"github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/orm/lib/migrate"
)

func main() {
	fmt.Print("Start")

	ormConfig := config.ORMConfig{
		DbDriver: "postgres",
		Host:     "127.0.0.1",
		Port:     5432,
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
}
