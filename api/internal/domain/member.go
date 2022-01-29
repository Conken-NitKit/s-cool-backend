package domain

import (
	"condog/pkg/utils/cerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Member struct {
	userId                   ID
	communityMemberId        ID
	communityMemberDisplayId ID
	role                     MemberRole
	banned                   bool
}

func (member Member) Validate() error {
	return cerror.TogetherError(
		member.userId.Validate(),
		member.communityMemberDisplayId.Validate(),
		member.communityMemberId.Validate(),
		member.role.Validate(),
	)
}

type MemberRole string

func (role MemberRole) Validate() error {
	return validation.Validate(
		string(role),
		validation.In("owner", "admin", "member"),
	)
}
