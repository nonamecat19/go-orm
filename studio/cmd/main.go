package main

import (
	adapterpostgres "adapter-postgres"
	"context"
	_ "github.com/lib/pq"
	appEntities "github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/studio/internal/app"
	config2 "github.com/nonamecat19/go-orm/studio/internal/config"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	postgresConfig := config.ORMConfig{
		Host:     "127.0.0.1",
		Port:     15432,
		User:     "postgres",
		Password: "root",
		DbName:   "orm",
		SSLMode:  false,
	}

	postgresAdapter := adapterpostgres.AdapterPostgres{}
	client := client2.CreateClient(postgresConfig, postgresAdapter)

	tables := []coreEntities.IEntity{
		appEntities.Order{},
		appEntities.Product{},
		appEntities.Role{},
		appEntities.User{},
	}

	cfg := config2.NewConfig()

	if err := app.Run(ctx, tables, client, cfg); err != nil {
		return err
	}

	return nil
}
