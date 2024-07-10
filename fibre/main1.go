package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the homepage!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About us page")
	})

	app.Listen(":3000")
}
