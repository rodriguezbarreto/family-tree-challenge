package structerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidatorStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())

	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")

	case "min":
		return errors.New(field + " is required with min " + validationError.Param())
	}

	return nil

}
