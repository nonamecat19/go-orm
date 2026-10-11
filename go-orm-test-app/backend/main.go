package main

import (
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	"log"
	"shopping-list/backend/database"
	"shopping-list/backend/entities"
	"shopping-list/backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	studio "github.com/nonamecat19/go-orm/studio/lib/app"
)

func getRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return app
}

func main() {
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	database.InitDbClient()
	app := getRouter()

	setupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	lists := api.Group("/lists")
	lists.Post("/", handlers.CreateList)
	lists.Get("/", handlers.GetLists)
	lists.Get("/:id", handlers.GetList)
	lists.Delete("/:id", handlers.DeleteList)

	listItems := lists.Group("/:listId/items")
	listItems.Post("/", handlers.AddItemToList)
	listItems.Delete("/:itemId", handlers.RemoveItemFromList)

	items := api.Group("/items")
	items.Post("/", handlers.CreateItem)
	items.Get("/", handlers.GetItems)
	items.Patch("/:id", handlers.UpdateItem)
	items.Delete("/:id", handlers.DeleteItem)

	tables := []coreEntities.Entity{
		entities.Item{},
		entities.List{},
	}

	studio.AddStudioFiberGroup(app, tables, database.DbClient, "/api/admin")
}
