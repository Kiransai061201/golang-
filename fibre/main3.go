package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", "Kiran")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		user := c.Locals("user").(string)
		return c.SendString("Hello, " + user)
	})

	app.Listen(":3000")
}
