package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"site/database"
)

func main() {
	// Start a new fiber app
	engine := html.New("../views", ".html")
	database.ConnectDB()
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello",
		})
	})
	app.Post("/", func(c *fiber.Ctx) error {

		var body struct {
			Message string
		}
		if err := c.BodyParser(&body); err != nil {
			fmt.Println("!!!")
			return err

		}
		fmt.Println(body.Message)
		return c.Render("index", fiber.Map{
			"Title":   "Hello world" + body.Message,
			"Message": body.Message,
		})
	})
	// Listen on PORT 300
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
