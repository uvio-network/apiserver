package converter

import (
	"strconv"
	"time"
)

func StringToTime(str string) time.Time {
	sec, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Unix(sec, 0).UTC()
}

func TimeToString(tim time.Time) string {
	if !tim.IsZero() {
		return strconv.FormatInt(tim.Unix(), 10)
	}

	return ""
}
