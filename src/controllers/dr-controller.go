package controllers

import "github.com/gofiber/fiber/v2"

type DeliveryReport struct {
}

func HandlerDeliveryReport(c *fiber.Ctx) error {

	// msisdn := c.Query("msisdn")
	// status := c.Query("status")
	// message := c.Query("message")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
