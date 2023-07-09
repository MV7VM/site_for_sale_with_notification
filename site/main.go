package main

import (
	"github.com/gofiber/fiber/v2"
	"site/database"
)

func main() {
	// Start a new fiber app
	app := fiber.New()
	database.ConnectDB()
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})
	// Listen on PORT 300
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
