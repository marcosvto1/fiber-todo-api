package tags

import "github.com/gofiber/fiber/v2"

func SetRoute(r fiber.Router) {
	tags := r.Group("/tags")
	tags.Post("/", addTag)
	tags.Get("/", getAll)
	tags.Get("/:id", getById)
	tags.Delete("/:id", deleteTag)
	tags.Put("/:id", updateTag)
}
