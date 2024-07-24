package converter

import "strconv"

func StringToFloat(str string) float64 {
	flo, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return flo
}

func FloatToString(flo float64) string {
	return strconv.FormatFloat(flo, 'f', -1, 64)
}
