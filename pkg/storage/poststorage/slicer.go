package poststorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Slicer []*Object

func (s Slicer) ID() []objectid.ID {
	var lis []objectid.ID

	for _, x := range s {
		lis = append(lis, x.ID)
	}

	return lis
}

func (s Slicer) ObjectID(cla objectid.ID) []*Object {
	var lis []*Object

	for _, x := range s {
		if cla == x.ID {
			lis = append(lis, x)
		}
	}

	return lis
}

func (s Slicer) Parent() []objectid.ID {
	var lis []objectid.ID

	for _, x := range s {
		if x.Parent != "" {
			lis = append(lis, x.Parent)
		}
	}

	return lis
}
