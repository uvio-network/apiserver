package generic

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

func Arg1[T string | objectid.ID](str string, one []T) []string {
	return For(len(one), func(i int) string { return fmt.Sprintf(str, one[i]) })
}
