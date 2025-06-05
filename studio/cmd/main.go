package main

import (
	adapterpostgres "adapter-postgres/lib"
	"context"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	appEntities "github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	config2 "github.com/nonamecat19/go-orm/studio/internal/config"
	"github.com/nonamecat19/go-orm/studio/lib/app"
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
	_, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
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

	tables := []coreEntities.Entity{
		appEntities.Order{},
		appEntities.Role{},
		appEntities.User{},
	}

	cfg := config2.NewConfig()

	server := fiber.New(fiber.Config{})

	app.AddStudioFiberGroup(server, tables, client, "/studio")

	err := server.Listen(cfg.ServerAddr)
	if err != nil {
		return err
	}

	return nil
}
