package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func updateTag(c *fiber.Ctx) error {
	tagBodyData := new(TagEntity)
	if err := c.BodyParser(tagBodyData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	tagResult := TagEntity{}
	err := db.UpdateOne("tags", c.Params("id"), tagBodyData, &tagResult)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(tagResult)
}
