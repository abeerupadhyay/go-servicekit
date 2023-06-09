package validators

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// S represents a sample struct with all the reuired validators
// added to its fields. This helps identifying all validation errors
// through a single struct
type S struct {
	A string   `json:"a" validate:"required,ascii,len=5"`
	B *int     `json:"b" validate:"omitempty,min=5,max=10"`
	C *float64 `json:"c" validate:"omitempty,min=5.5,max=10.5"`
	D []string `json:"d" validate:"omitempty,min=2"`

	URL         string `json:"url" validate:"omitempty,url"`
	CountryCode string `json:"country_code" validate:"omitempty,iso3166_1_alpha2"`
	Language    string `json:"language" validate:"omitempty,language_iso6391"`
}

func TestStructValidation(t *testing.T) {

	testcases := []struct {
		Name      string
		Data      string
		FieldName string
		ErrorTag  string
	}{
		{
			Name:      "required tag error message",
			Data:      `{"a": ""}`,
			FieldName: "a",
			ErrorTag:  "required",
		},
		{
			Name:      "ascii tag error message",
			Data:      `{"a": "ƒßeta"}`,
			FieldName: "a",
			ErrorTag:  "ascii",
		},
		{
			Name:      "len tag error message",
			Data:      `{"a": "foo123"}`,
			FieldName: "a",
			ErrorTag:  "len:5",
		},
		{
			Name:      "min tag (int) error message",
			Data:      `{"a": "foo12", "b": 1}`,
			FieldName: "b",
			ErrorTag:  "min:5",
		},
		{
			Name:      "max tag (int) error message",
			Data:      `{"a": "foo12", "b": 222}`,
			FieldName: "b",
			ErrorTag:  "max:10",
		},
		{
			Name:      "min tag (float) error message",
			Data:      `{"a": "foo12", "c": 1.0}`,
			FieldName: "c",
			ErrorTag:  "min:5.5",
		},
		{
			Name:      "max tag (float) error message",
			Data:      `{"a": "foo12", "c": 22.2}`,
			FieldName: "c",
			ErrorTag:  "max:10.5",
		},
		{
			Name:      "min tag (slice) error message",
			Data:      `{"a": "foo12", "d": ["foo"]}`,
			FieldName: "d",
			ErrorTag:  "min:2",
		},
		{
			Name:      "url tag error message",
			Data:      `{"a": "foo12", "url": "/path/to/resource"}`,
			FieldName: "url",
			ErrorTag:  "url",
		},
		{
			Name:      "language tag error message",
			Data:      `{"a": "foo12", "language": "qq"}`,
			FieldName: "language",
			ErrorTag:  "language_iso6391",
		},
		{
			Name:      "country_code tag error message",
			Data:      `{"a": "foo12", "country_code": "in"}`,
			FieldName: "country_code",
			ErrorTag:  "iso3166_1_alpha2",
		},
	}

	errMsg := "validation failed for tag '%s'"

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {

			assert := assert.New(t)

			s := &S{}
			_ = json.Unmarshal([]byte(tc.Data), s)

			err := Struct(s)
			assert.NotNil(err)

			valErr, ok := err.(*ValidationError)
			assert.True(ok)
			assert.Equal(tc.FieldName, valErr.Field, tc.Data)
			assert.Equal(fmt.Sprintf(errMsg, tc.ErrorTag), valErr.Description, tc.Data)
		})
	}
}
