package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"shopping-list/backend/entities"
	"shopping-list/backend/services"
	"strconv"
)

func CreateItem(c *fiber.Ctx) error {
	var itemCreate services.ItemCreate
	if err := c.BodyParser(&itemCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Непрвильний тип даних",
		})
	}

	if itemCreate.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Назва товару обов'язкова",
		})
	}

	newItem := entities.Item{
		Name:   itemCreate.Name,
		Bought: false,
	}

	err := services.CreateItem(newItem)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Не вдалось додати товар: %s", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Товар створено успішно",
	})
}

func GetItems(c *fiber.Ctx) error {
	err, items := services.GetAllItems()

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Не вдалось отримати товари: %s", err),
		})
	}

	return c.JSON(items)
}

func UpdateItem(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	var itemUpdate services.ItemUpdate
	if err = c.BodyParser(&itemUpdate); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = services.UpdateItem(itemUpdate, id)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update item",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Товар оновлено успішно",
	})
}

func DeleteItem(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	err = services.DeleteItem(id)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Товар видалено успішно",
	})
}
