package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func (controller *TagController) addTag(c *fiber.Ctx) error {
	tag := new(TagEntity)

	if err := c.BodyParser(tag); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}

	objectId, err := db.Insert(controller.Collection, tag)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}

	tag.ID = objectId

	return c.Status(http.StatusCreated).JSON(tag)
}
