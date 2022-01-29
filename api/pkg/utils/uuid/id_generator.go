//go:build !faketime
// +build !faketime

package uuid

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuidObj, _ := uuid.NewUUID()
	data := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
	uuidObj2 := uuid.NewSHA1(uuidObj, data)
	return strings.ReplaceAll(uuidObj2.String(), "-", "")[:15]
}
