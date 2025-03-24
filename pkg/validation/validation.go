package validation

import (
	"regexp"

	Validator "github.com/go-playground/validator/v10"
)

func NewValidator() *Validator.Validate {
	validator := Validator.New(Validator.WithRequiredStructEnabled())
	err := validator.RegisterValidation("fullname", fullnameValidation)
	if err != nil {
		panic(err)
	}
	return validator
}

func fullnameValidation(fullname Validator.FieldLevel) bool {
	matched, err := regexp.Match("[a-zA-Z]", []byte(fullname.Field().String()))
	if !matched || err != nil {
		return false
	}

	matched, err = regexp.Match("^[a-zA-Z\\-\\.\\s]+$", []byte(fullname.Field().String()))
	if !matched || err != nil {
		return false
	}

	return true
}
