package domain

import (
	"condog/pkg/utils/cerror"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	userId 					ID
	accountId			 	ID
	email 					Email
	name  					DisplayName
	gender					Gender
	birthDay				time.Time
	privacyStatus		PrivacyStatus
	timestamp				Timestamp
}

func (user User) Validate() error {
	return cerror.TogetherError(
		user.userId.Validate(),
		user.accountId.Validate(),
		user.email.Validate(),
		user.name.Validate(),
		user.gender.Validate(),
		user.privacyStatus.Validate(),
		user.timestamp.Validate(),
	)
}

type DisplayName string

func (name DisplayName) Validate() error {
	return validation.Validate(
		string(name),
		validation.Required.Error("DisplayName is required"),
		validation.RuneLength(1, 50),
	)
}

type Email string

func (email Email) Validate() error {
	return validation.Validate(
		string(email),
		validation.When(email != "", is.Email),
	)
}

type Gender string

func (geneder Gender) Validate() error {
	return validation.Validate(
		string(geneder),
		validation.In("Female", "Male", "Other"),
	)
}

type PrivacyStatus string

func (status PrivacyStatus) Validate() error {
	return validation.Validate(
		string(status),
		validation.In("public", "private"),
	)
}
