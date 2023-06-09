package validators

import (
	"testing"
)

func AssertValidationError(t *testing.T, err error, field string) {
	verr, ok := err.(*ValidationError)
	if !ok {
		t.Error("Unable to type assert ValidationError")
	}
	if verr == nil {
		t.Error("No error raised")
	} else if field != verr.Field {
		t.Errorf("Validation failed for field '%s' instead of '%s'", verr.Field, field)
	}
}
