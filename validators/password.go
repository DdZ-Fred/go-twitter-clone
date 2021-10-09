package validators

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func Password(fl validator.FieldLevel) bool {
	length := len(fl.Field().String())
	lower := 0
	upper := 0
	number := 0
	special := 0

	if length < 8 || length > 30 {
		return false
	}

	for _, c := range fl.Field().String() {
		switch {
		case unicode.IsNumber(c):
			number += 1
		case unicode.IsUpper(c):
			upper += 1
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special += 1
		case unicode.IsLetter(c) || c == ' ':
			lower += 1
		default:
			fmt.Println("Unkown character type: %s", c)
		}
	}
	return lower >= 1 && upper >= 1 && number >= 1 && special >= 1
}
