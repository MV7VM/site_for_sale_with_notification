package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
	"site/database"
	"site/internal/model"
	"time"
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
		var body1 model.NoteTask
		if err := c.BodyParser(&body); err != nil {
			fmt.Println("!!!")
			return err
		}
		if err := c.BodyParser(&body1); err != nil {
			fmt.Println("!!!")
			return err
		}
		now := time.Now()
		body1.TodayDate = now.Format("02.01.2006")
		body.ID = uuid.New()
		body1.ID = uuid.New()
		fmt.Println(body1.TodayDate)
		fmt.Println(body.ID, body.WhatMaterial1, body.Value1, body.Address, body.Date, body.PhoneNumber, body.Name, body.Comment, body1.TodayDate)
		err := database.DB.Create(&body).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
		}
		err = database.DB.Create(&body1).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
		}
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
