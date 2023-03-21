package goutil

import (
	"crypto/sha256"
)

func Hash(text string, salt string, salt2 string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt + text + salt2))

	return string(hasher.Sum(nil))
}
