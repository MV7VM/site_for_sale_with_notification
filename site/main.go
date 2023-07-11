package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"site/database"
	"site/internal/model"
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

		var body model.Note
		if err := c.BodyParser(&body); err != nil {
			fmt.Println("!!!")
			return err

		}
		fmt.Println(body.ID, body.WhatMaterial1, body.Value1, body.Address, body.Date, body.PhoneNumber, body.Comment)
		return c.Render("index", fiber.Map{
			"Title":   "Hello world" + body.WhatMaterial1,
			"Message": body.Value1,
		})
	})
	// Listen on PORT 300
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
