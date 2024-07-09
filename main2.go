package main

import (
	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	println("Request received:", c.Path())
	return c.Next()
}

func main() {
	app := fiber.New()

	app.Use(Logger)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Listen(":3000")
}
