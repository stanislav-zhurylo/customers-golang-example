package utils

import "github.com/go-playground/validator"

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Wrong email format"
	}
	return "Unknown error"
}
