package domain

import (
	"condog/pkg/utils/cerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ExternalNames struct {
	line    string
	discord string
}

type ExternalIds struct {
	twitter     string
	instagram   string
	note        string
	github      string
	atCoder     string
	qiita       string
	zenn        string
	speakerDeck string
}

type ExternalInfoNecessaryLevel struct {
	lineName      NecessaryLevel
	discordName   NecessaryLevel
	twitterId     NecessaryLevel
	instagramId   NecessaryLevel
	noteId        NecessaryLevel
	githubId      NecessaryLevel
	atCoderId     NecessaryLevel
	qiitaId       NecessaryLevel
	zennId        NecessaryLevel
	speakerDeckId NecessaryLevel
}

func (info ExternalInfoNecessaryLevel) Validate() error {
	return cerror.TogetherError(
		info.lineName.Validate(),
		info.discordName.Validate(),
		info.twitterId.Validate(),
		info.instagramId.Validate(),
		info.noteId.Validate(),
		info.githubId.Validate(),
		info.atCoderId.Validate(),
		info.qiitaId.Validate(),
		info.zennId.Validate(),
		info.speakerDeckId.Validate(),
	)
}

type NecessaryLevel string

const (
	NECESSARY_MUST        NecessaryLevel = "must"
	NECESSARY_WANT        NecessaryLevel = "want"
	NECESSARY_UNNECESSARY NecessaryLevel = ""
)

func (level NecessaryLevel) Validate() error {
	return validation.Validate(
		string(level),
		validation.Required,
		validation.In(string(NECESSARY_MUST), string(NECESSARY_WANT), string(NECESSARY_UNNECESSARY)),
	)
}
