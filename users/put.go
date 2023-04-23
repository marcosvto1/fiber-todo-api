package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
)

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(UserEntity)
	var userUpdated UserEntity

	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid json")
	}

	err := db.UpdateOne("users", id, user, &userUpdated)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(userUpdated)
}
