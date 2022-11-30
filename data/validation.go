package data

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Validate(i interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(i)
}

// validateSKU
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}
