package main

import (
	adapterpostgres "adapter-postgres/lib"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nonamecat19/go-orm/app/entities"
	"github.com/nonamecat19/go-orm/core/lib/config"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	querybuilder "github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	studioApp "github.com/nonamecat19/go-orm/studio/lib/app"
)

func main() {
	server := fiber.New()
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

	server.Get("/api/users", func(c *fiber.Ctx) error {
		var users []entities.User
		err := querybuilder.CreateQueryBuilder(client).
			Preload("orders").
			Preload("role").
			FindMany(&users)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(users)
	})

	server.Get("/api/users/:id", func(c *fiber.Ctx) error {
		var user entities.User
		err := querybuilder.CreateQueryBuilder(client).
			Where("id = ?", c.Params("id")).
			Preload("orders").
			Preload("role").
			Limit(1).
			FindMany(&user)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(user)
	})

	tables := []coreEntities.IEntity{
		entities.Order{},
		entities.Role{},
		entities.User{},
	}

	studioApp.AddStudioFiberGroup(server, tables, client, "/studio")

	_ = server.Listen(":7777")
}
