package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func updateTask(c *fiber.Ctx) error {
	taskBodyData := new(TaskEntity)

	if err := c.BodyParser(taskBodyData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	var task = TaskEntity{}
	err := db.UpdateOne("tasks", c.Params("id"), taskBodyData, &task)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    fiber.StatusBadGateway,
		})
	}

	return c.JSON(task)
}
