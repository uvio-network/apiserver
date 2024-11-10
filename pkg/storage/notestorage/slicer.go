package notestorage

import "slices"

type Slicer []*Object

// --------------------------------------------------------------------- //

func (s Slicer) Kind() []string {
	var lis []string

	see := map[string]struct{}{}
	for _, x := range s {
		if x.Kind == "" {
			continue
		}

		var exi bool
		{
			_, exi = see[x.Kind]
		}

		if exi {
			continue
		}

		{
			lis = append(lis, x.Kind)
		}

		{
			see[x.Kind] = struct{}{}
		}
	}

	return lis
}

// --------------------------------------------------------------------- //

func (s Slicer) ObjectKind(kin []string) Slicer {
	if len(kin) == 1 && kin[0] == "*" {
		return s
	}

	var lis Slicer

	for _, x := range s {
		if slices.Contains(kin, x.Kind) {
			lis = append(lis, x)
		}
	}

	return lis
}
