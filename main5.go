package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Page not found")
	})

	app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	})

	app.Listen(":3000")
}
