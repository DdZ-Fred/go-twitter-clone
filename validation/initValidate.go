package validation

import (
	"github.com/DdZ-Fred/go-twitter-clone/validators"
	"github.com/go-playground/validator/v10"
)

func InitValidate() *validator.Validate {
	validate := validator.New()

	// Add all custom validators here
	validate.RegisterValidation("password", validators.Password)

	return validate
}
