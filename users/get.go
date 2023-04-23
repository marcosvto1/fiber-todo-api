package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func getAll(c *fiber.Ctx) error {
	var users = []UserEntity{}

	err := db.Find("users", &users)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(users)
}

func getById(c *fiber.Ctx) error {
	var user UserEntity

	id := c.Params("id")

	err := db.FindById("users", id, &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}
