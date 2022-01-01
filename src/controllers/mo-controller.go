package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type MessageOriginated struct {
	Msisdn    string `validate:"required"`
	ShortCode string `validate:"required"`
	message   string `validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(mo MessageOriginated) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(mo)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func IncomeMessageOriginated(c *fiber.Ctx) error {

	mo := new(MessageOriginated)

	// if msisdn := c.Query("mobile_no"); msisdn != "" {

	// }
	// if shortcode := c.Query("short_code"); shortcode != "" {

	// }
	// if message := c.Query("message"); message != "" {

	// }

	errors := ValidateStruct(*mo)

	if errors != nil {
		return c.JSON(errors)

	}

	return c.JSON(fiber.Map{"message": "ok"})
}
