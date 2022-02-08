package domain

import validation "github.com/go-ozzo/ozzo-validation"

type School struct {
	schoolId ID
	name     Name
}

type SchoolMember struct {
	memberId        ID
	userId          ID
	circleId        ID
	mainDisplayName OptionalName
	subDisplayName  OptionalName
	role            SchoolRole
}

type SchoolRole string

const (
	SCHOOL_ROLE_OWNER SchoolRole = "Owner"
	SCHOOL_ROLE_ADMIN SchoolRole = "admin"
	SCHOOL_ROLE_NONE  SchoolRole = ""
)

func (role CircleRole) SchoolRole() error {
	return validation.Validate(
		string(role),
		validation.In(
			string(SCHOOL_ROLE_OWNER),
			string(SCHOOL_ROLE_ADMIN),
			string(SCHOOL_ROLE_NONE),
		),
	)
}
