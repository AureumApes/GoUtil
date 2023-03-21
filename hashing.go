package goutil

import (
	"crypto/sha256"
)

func Hash(text string, salt string, salt2 string) string {
	return string(sha256.New().Sum([]byte(salt + text + salt2)))
}
