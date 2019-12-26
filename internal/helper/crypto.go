package helper

import (
	"crypto/sha512"
	"fmt"
)

// GenerateSHA512 create sha256 by text
func GenerateSHA512(text string) string {
	data := []byte(text)
	hash := sha512.Sum512(data)
	return fmt.Sprintf("%x", hash[:])
}
