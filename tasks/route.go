package tasks

import "github.com/gofiber/fiber/v2"

func SetRoute(r fiber.Router) {
	tasks := r.Group("/tasks")
	tasks.Get("/", getAllTask)
	tasks.Get("/", getById)
	tasks.Post("/", addTask)
	tasks.Put("/:id", updateTask)
	tasks.Delete("/:id", deleteTask)
}
