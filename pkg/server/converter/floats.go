package converter

import (
	"strconv"
	"strings"
)

func FloatsToString(lis []float64) string {
	var buf strings.Builder

	for i, x := range lis {
		if i > 0 {
			buf.WriteString(",")
		}

		{
			buf.WriteString(strconv.FormatFloat(x, 'f', -1, 64))
		}
	}

	return buf.String()
}
