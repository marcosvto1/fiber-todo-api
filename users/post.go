package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func addUser(c *fiber.Ctx) error {
	body := new(UserEntity)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	id, err := db.Insert("users", body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}

	body.ID = id

	return c.Status(http.StatusCreated).JSON(body)

}
