package domain

import (
	"condog/pkg/utils/cerror"
)

type Community struct {
	id                         ID
	name                       Name
	description                Description
	externalInfoNecessaryLevel ExternalInfoNecessaryLevel
	training                   []Training
	// デザインが完成してからドメインを定義する
	// infomations []Infomation
	// projects    []Projects
}

func (community Community) Validate() error {
	return cerror.TogetherError(
		community.id.Validate(),
		community.name.Validate(),
		community.externalInfoNecessaryLevel.Validate(),
	)
}

type LinkItem struct {
	id           ID
	title        Name
	description  Description
	externalLink URL
}

func (item LinkItem) Validate() error {
	return cerror.TogetherError(
		item.id.Validate(),
		item.title.Validate(),
		item.externalLink.Validate(),
	)
}
