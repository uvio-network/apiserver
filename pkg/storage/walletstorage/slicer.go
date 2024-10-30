package walletstorage

import "github.com/xh3b4sd/objectid"

type Slicer []*Object

// --------------------------------------------------------------------- //

func (s Slicer) Owner() []objectid.ID {
	var lis []objectid.ID

	see := map[objectid.ID]struct{}{}
	for _, x := range s {
		if x.Owner == "" {
			continue
		}

		var exi bool
		{
			_, exi = see[x.Owner]
		}

		if exi {
			continue
		}

		{
			lis = append(lis, x.Owner)
		}

		{
			see[x.Owner] = struct{}{}
		}
	}

	return lis
}

// --------------------------------------------------------------------- //

func (s Slicer) ObjectKind(kin string) Slicer {
	var lis Slicer

	for _, x := range s {
		if x.Kind == kin {
			lis = append(lis, x)
		}
	}

	return lis
}
