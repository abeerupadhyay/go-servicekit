package validators

import (
	iso6391 "github.com/emvi/iso-639-1"
	"github.com/go-playground/validator/v10"
)

func ValidateLanguageISO6391(fl validator.FieldLevel) bool {
	return iso6391.ValidCode(fl.Field().String())
}
