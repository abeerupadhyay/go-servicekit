package validators

import (
	"fmt"
)

const DefaultError = "invalid_value"

type ValidationError struct {
	Err         string `json:"error"`
	Description string `json:"description"`
	Field       string `json:"field_name,omitempty"`
}

func (v *ValidationError) Error() string {
	if v.Field != "" {
		return fmt.Sprintf("%s. %s: %s", v.Err, v.Field, v.Description)
	} else {
		return fmt.Sprintf("%s. %s", v.Err, v.Description)
	}
}

func NewValidationError(desc string) *ValidationError {
	return &ValidationError{
		Err:         DefaultError,
		Description: desc,
	}
}

func NewValidationErrorf(desc string, args ...any) *ValidationError {
	return NewValidationError(fmt.Sprintf(desc, args...))
}

func NewValidationErrorWithField(field, desc string) *ValidationError {
	err := NewValidationError(desc)
	err.Field = field
	return err
}

func NewValidationErrorfWithField(field, desc string, args ...any) *ValidationError {
	return NewValidationErrorWithField(field, fmt.Sprintf(desc, args...))
}
