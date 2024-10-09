package walletstorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Slicer []*Object

// --------------------------------------------------------------------- //

func (s Slicer) Owner() []objectid.ID {
	var lis []objectid.ID

	for _, x := range s {
		lis = append(lis, x.Owner)
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
