package walletaddress

import (
	"fmt"
	"regexp"
)

var (
	forExp = regexp.MustCompile(`^(0x[a-zA-Z0-9]{4})[a-zA-Z0-9]+([a-zA-Z0-9]{4})$`)
	verExp = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
)

func Format(str string) string {
	if str == "" {
		return ""
	}

	var mat []string
	{
		mat = forExp.FindStringSubmatch(str)
	}

	if mat == nil {
		return str
	}

	return fmt.Sprintf("%s•••%s", mat[1], mat[2])
}

func Verify(str string) bool {
	return verExp.MatchString(str)
}
