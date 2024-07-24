package converter

import "strconv"

func StringToBool(str string) bool {
	boo, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}

	return boo
}

func BoolToString(boo bool) string {
	return strconv.FormatBool(boo)
}
