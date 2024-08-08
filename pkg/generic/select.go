package generic

import "github.com/uvio-network/apiserver/pkg/object/objectid"

// Select returns the subset of elements that are not present in all.
func Select[T string | objectid.ID | int64](all []T, sub []T) []T {
	if len(sub) == 0 {
		return nil
	}

	m := make(map[T]struct{})
	for _, x := range all {
		m[x] = struct{}{}
	}

	var l []T
	for _, x := range sub {
		_, e := m[x]
		if !e {
			l = append(l, x)
		}
	}

	return l
}
