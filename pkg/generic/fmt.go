package generic

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

func Fmt[T string | objectid.ID](lis []T, str string) []string {
	return Func(lis, func(x T) string { return fmt.Sprintf(str, x) })
}
