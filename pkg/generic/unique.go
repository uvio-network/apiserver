package generic

import "github.com/uvio-network/apiserver/pkg/object/objectid"

// Unique returns the distinct items of the given list without preserving any
// particular input order.
func Unique[T string | objectid.ID | int64](lis []T) []T {
	see := map[T]struct{}{}

	for _, x := range lis {
		see[x] = struct{}{}
	}

	var uni []T

	for k := range see {
		uni = append(uni, k)
	}

	return uni
}
