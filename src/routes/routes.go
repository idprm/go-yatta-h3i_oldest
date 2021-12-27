package routes

import (
	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("v1")
	v1.Post("/admin/register", controllers.Register)
}
