package handlers

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/studio/internal/config"
	"github.com/nonamecat19/go-orm/studio/internal/view/settings"
	"github.com/nonamecat19/go-orm/studio/internal/view/tables"
)

func TablesPage(c *fiber.Ctx) error {
	//config.PostgresConfig

	dbClient := client.CreateClient(config.PostgresConfig)
	db := dbClient.GetDb()

	rows, err := db.Query(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public';
    `)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tableName string
	fmt.Println("Tables in the public schema:")
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}
		fmt.Println(tableName)
	}

	props := tables.TablePageProps{
		Tables: []tables.Table{
			{Title: "Users", ID: "1"},
			{Title: "Orders", ID: "2"},
			{Title: "Products", ID: "3"},
		},
	}

	return Render(c, tables.TablesPage(props))
}

func SettingsPage(c *fiber.Ctx) error {
	return Render(c, settings.SettingsPage())
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
