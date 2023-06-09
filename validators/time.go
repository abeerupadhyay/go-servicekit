package validators

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	Regex12HFormatTime = regexp.MustCompile("((0?)[0-9]|1[0-2]):[0-5][0-9]")
	Regex24HFormatTime = regexp.MustCompile("([0-1][0-9]|2[0-3]):[0-5][0-9]")
)

func Validate12HFormatTime(fl validator.FieldLevel) bool {
	t := fl.Field().String()
	return Regex12HFormatTime.MatchString(t)
}

func Validate24HFormatTime(fl validator.FieldLevel) bool {
	t := fl.Field().String()
	return Regex24HFormatTime.MatchString(t)
}

func ValidateTimeRFC3339(fl validator.FieldLevel) bool {
	t := fl.Field().String()
	_, err := time.Parse(time.RFC3339, t)
	return err == nil
}
