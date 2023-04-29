package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
	"github.com/marcosvto1/fiber-todo-api/tags"
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

	err = tags.AddTask(taskBodyData.ID.Hex(), taskBodyData.Tags)
	if err != nil {
		db.Delete("tasks", taskBodyData.ID.Hex())
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
		})
	}

	return c.Status(http.StatusCreated).JSON(taskBodyData)
}
