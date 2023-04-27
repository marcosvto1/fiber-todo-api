package tags

import "github.com/gofiber/fiber/v2"

func SetRoute(r fiber.Router) {
	controller := NewTagController()

	tags := r.Group("/tags")
	tags.Post("/", controller.addTag)
	tags.Get("/", controller.getAll)
	tags.Get("/:id", controller.getById)
	tags.Delete("/:id", controller.deleteTag)
	tags.Put("/:id", controller.updateTag)
}
