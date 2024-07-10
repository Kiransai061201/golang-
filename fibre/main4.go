package main

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name  string `json:"Kiran"`
	Email string `json:"Kiransai0612@gmail.com"`
}

func main() {
	app := fiber.New()

	app.Post("/user", func(c *fiber.Ctx) error {
		var person Person
		if err := c.BodyParser(&person); err != nil {
			return err
		}
		return c.JSON(person)
	})

	app.Listen(":3000")
}
