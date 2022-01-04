package controllers

import (
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

	// SET HEADER
	c.Accepts("text/plain")
	c.AcceptsCharsets("utf-8", "iso-8859-1")
	c.Set("Content-Type", "text/plain")

	// QUERY PARAM ON GET METHOD
	// msisdn := c.Query("mobile_no")
	// shortcode := c.Query("short_code")

	// GET MESSAGE (SMS) TO DECODE VALUE
	decodedMessage, _ := url.QueryUnescape(c.Query("message"))

	var kewordThree string
	// SPLIT VALUES
	subKey := strings.Split(decodedMessage, " ")

	// CONDITION IF KEYWORD IS EMPTY
	if len(subKey) < 3 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not valid message"})
	} else {
		// REG OR UNREG
		// keyWordOne = subKey[0]
		// // SERVICE
		// keywordTwo = subKey[1]
		// ADNET
		kewordThree = subKey[2]
	}

	// SELECT ON TABLE KEYWORD
	var keyword models.Keyword
	database.Database.Db.Where("name = ?", kewordThree).First(&keyword)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": keyword.Adnet})
}
