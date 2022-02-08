package domain

import validation "github.com/go-ozzo/ozzo-validation"

type Circle struct {
	circleId ID
	schoolId ID
	Name     Name
}

type CircleMember struct {
	memberId ID
	userId   ID
	circleId ID
	role     CircleRole
}

type CircleRole string

const (
	CIRCLE_ROLE_OWNER CircleRole = "Owner"
	CIRCLE_ROLE_ADMIN CircleRole = "admin"
	CIRCLE_ROLE_NONE  CircleRole = ""
)

func (role CircleRole) Validate() error {
	return validation.Validate(
		string(role),
		validation.In(
			string(CIRCLE_ROLE_OWNER),
			string(CIRCLE_ROLE_ADMIN),
			string(CIRCLE_ROLE_NONE),
		),
	)
}
