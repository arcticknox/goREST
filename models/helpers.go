package models

import (
	"github.com/go-playground/validator/v10"
)

func GetErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "gt":
		return "the value of " + err.Field() + " must be greater than " + err.Param()
	case "gte":
		return "the value of " + err.Field() + " must be greater than or equals " + err.Param()
	default:
		return "validation error in " + err.Field()
	}

}
