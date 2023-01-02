package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ValidatorRequest(data any) []string {
	validate := validator.New()

	var msgErr []string
	if err := validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msgErr = append(msgErr, "Error with field "+err.Field()+" where Tag "+err.Tag())
		}
	}
	return msgErr
}
