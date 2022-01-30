package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Name string

func (name Name) Validate() error {
	return validation.Validate(
		string(name),
		validation.Required,
		validation.RuneLength(1, 50),
	)
}

type OptionalName string

func (name OptionalName) Validate() error {
	return validation.Validate(
		string(name),
		validation.When(name != "", validation.RuneLength(1, 50)),
	)
}

type Description string

type Email string

func (email Email) Validate() error {
	return validation.Validate(
		string(email),
		validation.When(email != "", is.Email),
	)
}

type URL string

func (url URL) Validate() error {
	return validation.Validate(
		string(url),
		is.URL,
	)
}
