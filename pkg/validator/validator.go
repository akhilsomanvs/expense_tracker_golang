package validator

import (
	"fmt"
	"strings"

	playground "github.com/go-playground/validator/v10"
)

var validate *playground.Validate

func init() {
	validate = playground.New()
}

func Validate(data any) error {
	if err := validate.Struct(data); err != nil {

		validationErrors, ok := err.(playground.ValidationErrors)
		if !ok {
			return err
		}

		messages := make([]string, 0, len(validationErrors))

		for _, fieldErr := range validationErrors {
			messages = append(messages, buildMessage(fieldErr))
		}
		return fmt.Errorf("%s", strings.Join(messages, ", "))
	}
	return nil
}

func buildMessage(err playground.FieldError) string {

	switch err.Tag() {

	case "required":
		return fmt.Sprintf("%s is required", err.Field())

	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())

	case "min":
		return fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())

	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", err.Field(), err.Param())

	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
