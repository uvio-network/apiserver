package generic

// Unique returns the distinct items of the given list without preserving any
// particular input order.
func Unique[T comparable](lis []T) []T {
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
