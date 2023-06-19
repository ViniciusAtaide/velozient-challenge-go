package passwordcards

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniciusataide/velozient-challenge-go/domain"
)

var validate = validator.New()

func validateStruct[Validatable domain.Validatable](strct Validatable) []*domain.ErrorResponse {
	var errors []*domain.ErrorResponse

	err := validate.Struct(strct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element domain.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func bodyValidator[Validatable domain.Validatable](c *fiber.Ctx) error {
	dto := new(Validatable)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := validateStruct(*dto); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func pathValidator[Validatable domain.Validatable](c *fiber.Ctx) error {
	dto := new(Validatable)
	if err := c.ParamsParser(dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := validateStruct(*dto); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
