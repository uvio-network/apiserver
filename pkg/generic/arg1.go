package generic

import (
	"fmt"

	"github.com/xh3b4sd/objectid"
)

func Arg1[T string | objectid.ID](str string, one []T) []string {
	return For(len(one), func(i int) string { return fmt.Sprintf(str, one[i]) })
}
