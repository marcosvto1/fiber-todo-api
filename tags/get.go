package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func getAll(c *fiber.Ctx) error {
	tags := []TagEntity{}

	err := db.Find("tags", &tags)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(tags)
}

func getById(c *fiber.Ctx) error {
	tag := TagEntity{}

	err := db.FindById("tags", c.Params("id"), tag)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(tag)
}
