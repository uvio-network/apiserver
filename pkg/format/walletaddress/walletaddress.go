package walletaddress

import (
	"regexp"
)

var (
	regExp = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
)

func Verify(str string) bool {
	return regExp.MatchString(str)
}
