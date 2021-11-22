package util

import "github.com/go-playground/validator/v10"

func ValidateRequiredFields(s interface{}) error {
	return validator.New().Struct(s)
}
