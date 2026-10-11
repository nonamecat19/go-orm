package handlers

import (
	"fmt"
	"shopping-list/backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddItemToList(c *fiber.Ctx) error {
	var addItemToList services.AddItemToListBody

	if err := c.BodyParser(&addItemToList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fmt.Println(addItemToList)

	err := services.AddItemToList(addItemToList)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Помилка додавання продукту до списку",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Продукт доданий у список успішно",
	})
}

func RemoveItemFromList(c *fiber.Ctx) error {
	listID, err := strconv.ParseInt(c.Params("listId"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Непрвильний тип даних",
		})
	}
	itemID, err := strconv.ParseInt(c.Params("itemId"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Непрвильний тип даних",
		})
	}

	err = services.RemoveItemFromList(itemID, listID)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Помилка видалення продукту зі списку",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Продукт видалений зі списку успішно",
	})
}
