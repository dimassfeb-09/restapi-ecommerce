package helpers

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidatorRequest(data any) []string {
	validate := validator.New()

	var msgErr []string
	if err := validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			tag := err.Tag()
			msgErr = append(msgErr, "Error with field "+field+" where must "+tag)
		}
	}
	return msgErr
}
