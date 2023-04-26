package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func addTag(c *fiber.Ctx) error {
	tag := new(TagEntity)

	if err := c.BodyParser(tag); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "invalid json",
		})
	}

	objectId, err := db.Insert("tags", tag)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	tag.ID = objectId

	return c.Status(http.StatusCreated).JSON(tag)
}
