package validators

import (
	"fmt"

	"github.com/abeerupadhyay/go-servicekit/str"
	"github.com/go-playground/validator/v10"
)

var V *validator.Validate

func init() {
	V = validator.New()

	// Custom time-related validators
	V.RegisterValidation("12htime", Validate12HFormatTime)
	V.RegisterValidation("24htime", Validate24HFormatTime)
	V.RegisterValidation("time_rfc3339", ValidateTimeRFC3339)

	// Language code validator
	V.RegisterValidation("language_iso6391", ValidateLanguageISO6391)
}

func Struct(v any) error {
	err := V.Struct(v)
	if err != nil {
		errors, _ := err.(validator.ValidationErrors)
		fe := errors[0]

		tag := fe.ActualTag()
		if fe.Param() != "" {
			tag = fmt.Sprintf("%s:%s", tag, fe.Param())
		}
		return NewValidationErrorfWithField(
			str.ToSnakeCase(fe.Field()), "validation failed for tag '%s'", tag)
	}

	return nil
}
