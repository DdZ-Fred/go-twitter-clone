package validation

import (
	"reflect"
	"strings"

	"github.com/DdZ-Fred/go-twitter-clone/validators"
	"github.com/go-playground/validator/v10"
)

func InitValidate() *validator.Validate {
	validate := validator.New()

	// Add all custom validators here
	validate.RegisterValidation("password", validators.Password)

	// Registers a function to get alternate names for StructFields
	// Returns the json property name rather than the actual struct property name
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}
