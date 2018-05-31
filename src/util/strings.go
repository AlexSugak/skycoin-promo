package util

import (
	"fmt"
	"strings"
)

// TrimLong accepts string add cut it if it's longer than maxLength
func TrimLong(str string, maxLength int) string {
	if len(str) < maxLength {
		return str
	}

	return str[0:1000]
}

// FirstToLower transforms first char of a string to lowercase
func FirstToLower(str string) string {
	if len(str) == 1 {
		return strings.ToLower(str)
	} else if len(str) > 1 {
		head := strings.ToLower(str[0:1])
		tail := str[1:]
		return fmt.Sprintf("%s%s", head, tail)
	}

	return str
}
