package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		for {
			msgType, msg, err := c.ReadMessage()
			if err != nil {
				return
			}
			if err := c.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	}))

	app.Listen(":3000")
}
