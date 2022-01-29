package domain

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ID string

func (id ID) Validate() error {
	return validation.Validate(
		string(id),
		validation.Required,
		validation.Length(4, 32),
		validation.Match(regexp.MustCompile("^[0-9a-z]*$")),
	)
}
