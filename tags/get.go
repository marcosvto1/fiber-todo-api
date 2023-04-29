package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func (controller *TagController) getAll(c *fiber.Ctx) error {
	tags := []TagEntity{}

	err := db.Find(controller.Collection, nil, &tags)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
		})
	}

	return c.JSON(tags)
}

func (controller *TagController) getById(c *fiber.Ctx) error {
	tag := TagEntity{}

	err := db.FindById(controller.Collection, c.Params("id"), tag)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusNotFound,
		})
	}

	return c.JSON(tag)
}
