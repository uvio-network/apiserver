package generic

import "github.com/uvio-network/apiserver/pkg/object/objectid"

// Duplicate expresses whether the given list contains duplicated items.
func Duplicate[T string | objectid.ID | int | int64](lis []T) bool {
	see := map[T]struct{}{}

	for _, x := range lis {
		{
			_, exi := see[x]
			if exi {
				return true
			}
		}

		{
			see[x] = struct{}{}
		}
	}

	return false
}
