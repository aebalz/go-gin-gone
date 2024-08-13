package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct using the validator library
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
