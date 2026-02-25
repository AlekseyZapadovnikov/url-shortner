package validation

import "github.com/go-playground/validator/v10"

var Valid *validator.Validate

func init() {
	Valid = initValidator()
}

func initValidator() *validator.Validate {
	v := validator.New()
	return v
}
