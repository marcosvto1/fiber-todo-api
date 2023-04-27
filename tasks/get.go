package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func getAllTask(c *fiber.Ctx) error {
	tasks := []TaskEntity{}

	err := db.Find("tasks", &tasks)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadGateway,
		})
	}

	return c.JSON(tasks)
}

func getById(c *fiber.Ctx) error {
	task := TaskEntity{}

	err := db.FindById("tasks", c.Params("id"), &task)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadGateway,
		})
	}

	return c.JSON(task)
}
