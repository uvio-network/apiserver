package generic

// Select returns the subset of elements that are not present in all.
func Select[T comparable](all []T, sub []T) []T {
	if len(sub) == 0 {
		return nil
	}

	m := make(map[T]struct{})
	for _, x := range all {
		m[x] = struct{}{}
	}

	s := make(map[T]struct{})

	var l []T
	for _, x := range sub {
		{
			_, e := s[x]
			if e {
				continue
			}
		}

		// Ensure that we return a clean list, even if sub contains duplicates.
		{
			s[x] = struct{}{}
		}

		{
			_, e := m[x]
			if !e {
				l = append(l, x)
			}
		}
	}

	return l
}
