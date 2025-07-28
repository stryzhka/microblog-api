package util

import (
	"crypto/sha1"
	"fmt"
)

func GeneratePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
