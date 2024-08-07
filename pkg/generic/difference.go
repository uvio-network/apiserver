package generic

import "github.com/uvio-network/apiserver/pkg/object/objectid"

// Difference returns the symmetric difference of the given slices.
func Difference[T string | objectid.ID | int64](a []T, b []T) []T {
	m := make(map[T]bool, len(a)+len(b))

	for _, x := range a {
		m[x] = true
	}

	for _, x := range b {
		if m[x] {
			delete(m, x)
		} else {
			m[x] = true
		}
	}

	l := make([]T, 0, len(m))

	for x := range m {
		l = append(l, x)
	}

	return l
}
