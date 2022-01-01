package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/routes"
)

func init() {
	// load config
	config, err := config.SetupConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to database
	database.Connect(config.DBSource)

	// connect to redis
	database.SetupRedis()

}

func main() {

	// declare fiber
	app := fiber.New()

	// setup routes
	routes.Setup(app)

	// start server on http
	log.Fatal(app.Listen(":8000"))
}
