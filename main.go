package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/database"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Get request")
	})

	app.Get("/:param", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Post request")
	})

	log.Fatal(app.Listen(":8000"))
}
