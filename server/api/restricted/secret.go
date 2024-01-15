package restricted

import (
	"math/rand"
	"strings"
)

const CHARSET = "1234567890abcdefghijklmnopqrstuvwxyz"

func GenerateSecret() string {
	secret := strings.Builder{}
	secret.Grow(64)

	for i := 0; i < 64; i++ {
		secret.WriteByte(CHARSET[rand.Intn(len(CHARSET))])
	}

	return secret.String()
}
