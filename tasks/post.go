package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func addTask(c *fiber.Ctx) error {
	taskBodyData := new(TaskEntity)
	if err := c.BodyParser(taskBodyData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	objectId, err := db.Insert("tasks", taskBodyData)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	taskBodyData.ID = objectId

	return c.Status(http.StatusCreated).JSON(taskBodyData)
}
