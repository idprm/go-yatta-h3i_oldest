package controllers

import "github.com/gofiber/fiber/v2"

type DeliveryReport struct {
}

func IncomeDeliveryReport(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
