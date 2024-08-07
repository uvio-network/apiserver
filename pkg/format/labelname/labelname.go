package labelname

import (
	"regexp"
	"strings"
)

var (
	labels = regexp.MustCompile(`^[A-Za-z0-9\-\s]+$`)
	spaces = regexp.MustCompile(`\s+`)
)

// Format cleans strings for their use as part of storage keys. For instance, we
// use Format for label names and want to ensure that letter casing does not
// define labels. The same word, but with a different capitalization should not
// make for a new label, because this behaviour would cause separate categories
// for claims. Thus "MEV" is indexed with "mev" and "DeFi" is indexed with
// "defi".
func Format(str string) string {
	// Replace multiple spaces with a single one.
	str = spaces.ReplaceAllString(str, " ")
	// Remove leading and trailing spaces.
	str = strings.TrimSpace(str)
	// Ensure only lower case letters.
	str = strings.ToLower(str)
	// Escape left over spaces.
	str = strings.ReplaceAll(str, " ", "-")

	return str
}

func Verify(lis []string) bool {
	for _, x := range lis {
		if !labels.MatchString(x) {
			return false
		}
	}

	return true
}
