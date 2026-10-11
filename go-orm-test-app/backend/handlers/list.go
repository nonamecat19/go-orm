package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"shopping-list/backend/database"
	"shopping-list/backend/entities"
	"shopping-list/backend/services"
	"strconv"
)

type ListCreate struct {
	Name string `json:"name"`
}

func CreateList(c *fiber.Ctx) error {
	var listCreate ListCreate
	if err := c.BodyParser(&listCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if listCreate.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "List name is required",
		})
	}

	newList := entities.List{
		Name: listCreate.Name,
	}

	err := querybuilder.
		CreateQueryBuilder(database.DbClient).
		InsertOne(newList)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Не вдалось додати список: %s", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Список створено успішно",
	})
}

func GetLists(c *fiber.Ctx) error {
	err, lists := services.GetAllLists()

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не вдалось отримати списки",
		})
	}

	return c.JSON(lists)
}

func GetList(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid list ID",
		})
	}

	err, listResult := services.GetListWithItems(id)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не вдалось отримати списки",
		})
	}

	return c.JSON(listResult)
}

func DeleteList(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	err = services.DeleteList(id)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete list",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "List deleted successfully",
	})
}
