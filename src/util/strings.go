package util

// TrimLong accepts string add cut it if it's longer than maxLength
func TrimLong(str string, maxLength int) string {
	if len(str) < maxLength {
		return str
	}

	return str[0:1000]
}
