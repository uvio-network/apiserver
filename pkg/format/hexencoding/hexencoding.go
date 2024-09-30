package hexencoding

import (
	"regexp"
)

var (
	regExp = regexp.MustCompile(`^0x[0-9a-fA-F]+$`)
)

func Verify(str string) bool {
	return regExp.MatchString(str) && !zerHsh(str)
}

func zerHsh(str string) bool {
	for _, char := range str[2:] {
		if char != '0' {
			return false
		}
	}

	return true
}
