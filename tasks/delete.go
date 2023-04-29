package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
	"github.com/marcosvto1/fiber-todo-api/tags"
)

func deleteTask(c *fiber.Ctx) error {
	err := tags.RemoveTasks(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadGateway,
		})
	}

	err = db.Delete("tasks", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadGateway,
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
