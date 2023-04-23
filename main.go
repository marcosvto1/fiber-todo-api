package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	v1 := app.Group("/api/v1")
	v1.Use(func(c *fiber.Ctx) error {
		c.Set("X-Frame-Options", "SAMEORIGIN")

		return c.Next()
	})
	v1.Get("/:name?", func(c *fiber.Ctx) error {
		s := c.Params("name")
		if s == "" {
			s = "World"
		}
		fmt.Fprintf(c, "Hello %s \n", s)
		return nil
	})

	app.Get("/user/+", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s", c.Params("+"))
		return nil
	})

	app.Get("/voos/:de-:para", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s-%s\n", c.Params("de"), c.Params("para"))
		return nil
	})

	app.Get("/planta/:tipo.:especie", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s.%s\n", c.Params("tipo"), c.Params("especie"))
		return nil
	})

	app.Listen(":8080")
}
