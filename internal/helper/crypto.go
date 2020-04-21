package helper

import (
	"crypto/md5"
	"fmt"
)

// MakeMD5String make md5 hash
func MakeMD5String(s string) string {
	stringByte := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(stringByte))
}
