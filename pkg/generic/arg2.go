package generic

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

func Arg2[T string | objectid.ID](str string, one []T, two []T) []string {
	return For(len(one), func(i int) string { return fmt.Sprintf(str, one[i], two[i]) })
}
