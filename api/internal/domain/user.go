package domain

import (
	"condog/pkg/utils/cerror"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	userId        ID
	accountId     ID
	email         Email
	name          DisplayName
	gender        Gender
	birthDay      time.Time
	privacyStatus PrivacyStatus
	timestamp     Timestamp
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
		validation.Required,
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

const (
	GENDER_FEMALE Gender = "female"
	GENDER_MALE   Gender = "male"
	GENDER_OTHER  Gender = ""
)

func (geneder Gender) Validate() error {
	return validation.Validate(
		string(geneder),
		validation.Required,
		validation.In(string(GENDER_FEMALE), string(GENDER_MALE), string(GENDER_OTHER)),
	)
}

type PrivacyStatus string

const (
	PRIVACY_PUBLIC  PrivacyStatus = "public"
	PRIVACY_PRIVATE PrivacyStatus = "private"
)

func (status PrivacyStatus) Validate() error {
	return validation.Validate(
		string(status),
		validation.Required,
		validation.In(string(PRIVACY_PUBLIC), string(PRIVACY_PRIVATE)),
	)
}
