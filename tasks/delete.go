package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func deleteTask(c *fiber.Ctx) error {
	err := db.Delete("tasks", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadGateway,
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
