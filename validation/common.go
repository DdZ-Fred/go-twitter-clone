package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	FailedField    string `json:"failedField"`
	ValidatorKey   string `json:"validatorKey"`
	ValidatorParam string `json:"validatorParam"`
}

func FailedValidationResponse(c *fiber.Ctx, validationErrors []*ValidationError) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"type":  "failed-payload-validation",
		"items": validationErrors,
	})
}

func ExtractValidationErrorsFromErr(validateStructErr error) []*ValidationError {
	var errors []*ValidationError

	if validateStructErr != nil {
		for _, err := range validateStructErr.(validator.ValidationErrors) {
			var element ValidationError
			element.FailedField = err.Field()
			element.ValidatorKey = err.Tag()
			element.ValidatorParam = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
