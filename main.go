package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/marcosvto1/fiber-todo-api/users"
)

func main() {
	app := fiber.New()

	// Set Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Create Versioning API
	v1 := app.Group("/api/v1")

	users.SetRoutes(v1)

	app.Listen(":8080")
}
