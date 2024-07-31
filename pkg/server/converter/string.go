package converter

import (
	"strings"
)

const (
	sep = ","
)

func StringToSlice(str string) []string {
	var lis []string

	for _, x := range strings.Split(str, sep) {
		lis = append(lis, strings.TrimSpace(x))
	}

	return lis
}

func SliceToString(lis []string) string {
	return strings.Join(lis, sep)
}
