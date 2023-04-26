package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func deleteTag(c *fiber.Ctx) error {
	err := db.Delete("tags", c.Params("id"))
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
