package domain

import (
	"time"
)

type User struct {
	accountId ID // ユーザーが定義するID
	userId    ID // 自動生成されるID
	email     Email
	name      Name
	birthDay  time.Time
	timestamp Timestamp
}
