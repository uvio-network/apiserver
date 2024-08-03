package votestorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) Claim() []objectid.ID {
	var lis []objectid.ID

	for _, x := range s {
		lis = append(lis, x.Claim)
	}

	return lis
}
