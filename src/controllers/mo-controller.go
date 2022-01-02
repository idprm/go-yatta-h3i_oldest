package controllers

import (
	"log"
	"net/url"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/models"
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

func HandlerMessageOriginated(c *fiber.Ctx) error {

	// mo := new(MessageOriginated)

	// if msisdn := c.Query("mobile_no"); msisdn != "" {

	// }
	// if shortcode := c.Query("short_code"); shortcode != "" {

	// }

	// GET MESSAGE (SMS)
	message := c.Query("message")

	// DECODE VALUE
	decodedValue, err := url.QueryUnescape(message)
	if err != nil {
		log.Fatal(err)
	}

	var word string
	// SPLIT VALUES
	subKey := strings.Split(decodedValue, " ")

	// CONDITION IF KEYWORD IS EMPTY
	if len(subKey) < 3 {
		return c.JSON(fiber.Map{"message": "keyword empty"})
	} else {
		word = subKey[2]
	}

	// errors := ValidateStruct(*mo)

	// if errors != nil {
	// 	return c.JSON(errors)
	// }

	// SELECT ON TABLE KEYWORD
	var keyword models.Keyword
	database.DB.Where("name = ?", word).First(&keyword)

	// if keyword.Id == 0 {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Invalid Credentials",
	// 	})
	// }

	return c.JSON(fiber.Map{"message": keyword})
}
