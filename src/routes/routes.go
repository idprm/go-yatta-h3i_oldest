package routes

import (
	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/controllers"
)

func Setup(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/v1")

	/******************************************
	**	API URL :
	**	GET /api/v1/moh3i
	**	GET /api/v1/drh3i
	*******************************************/
	v1.Get("/moh3i", controllers.HandlerMessageOriginated)
	v1.Get("/drh3i", controllers.HandlerDeliveryReport)
}
