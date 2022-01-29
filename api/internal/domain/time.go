package domain

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Timestamp struct {
	createdAt time.Time
	updatedAt time.Time
}

func (timestamp Timestamp) Validate() error {
	return validation.Validate(
		timestamp.ValidateOrder(),
	)
}

func (timestamp Timestamp) ValidateOrder() error {
	var createdAt = timestamp.createdAt.Unix()
	var updatedAt = timestamp.updatedAt.Unix()
	if createdAt > updatedAt {
		return errors.New("updatedAt cannot be set before the createdAt")
	}
	return nil
}
