package domain

import (
	"condog/pkg/utils/cerror"
)

type Trainee struct {
	communityMemberId ID
	completeLinks     []URL
}

type Training struct {
	title       Name
	description Description
	contents    []TrainingContent
}

type TrainingContent struct {
	title Name
	link  URL
}

func (content TrainingContent) Validate() error {
	return cerror.TogetherError(
		content.title.Validate(),
		content.link.Validate(),
	)
}
