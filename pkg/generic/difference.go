package generic

import "github.com/uvio-network/apiserver/pkg/object/objectid"

// Difference returns the symmetric difference of the given slices.
func Difference[T string | objectid.ID | int64](a []T, b []T) []T {
	m := make(map[T]bool, len(a)+len(b))

	for _, x := range a {
		m[x] = true
	}

	for _, x := range b {
		// We need to keep all elements in the map and check for key existance in
		// particular, order to cover potential duplicates.
		_, e := m[x]
		if e {
			m[x] = false
		} else {
			m[x] = true
		}
	}

	l := make([]T, 0, len(m))

	for x := range m {
		// Only keep the keys that are part of the difference.
		if m[x] {
			l = append(l, x)
		}
	}

	return l
}
