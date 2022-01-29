package domain

import (
	"condog/pkg/utils/cerror"
)

// 部内ユーザー向けのアカウント設定
type Student struct {
	User

	studentId			ID
}

func (student Student) Validate() error {
	return cerror.TogetherError(
		student.studentId.Validate(),
	)
}
