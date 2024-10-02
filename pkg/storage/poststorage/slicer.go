package poststorage

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
)

type Slicer []*Object

// --------------------------------------------------------------------- //

func (s Slicer) ID() []objectid.ID {
	var lis []objectid.ID

	for _, x := range s {
		lis = append(lis, x.ID)
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

// --------------------------------------------------------------------- //

func (s Slicer) ObjectID(oid []objectid.ID) Slicer {
	all := map[objectid.ID]struct{}{}
	for _, x := range oid {
		all[x] = struct{}{}
	}

	var lis Slicer

	for _, x := range s {
		var exi bool
		{
			_, exi = all[x.ID]
		}

		if exi {
			lis = append(lis, x)
		}
	}

	return lis
}

func (s Slicer) ObjectKind(kin string) Slicer {
	var lis Slicer

	for _, x := range s {
		if x.Kind == kin {
			lis = append(lis, x)
		}
	}

	return lis
}

func (s Slicer) ObjectLifecycle(lif objectlabel.DesiredLifecycle) Slicer {
	var lis Slicer

	for _, x := range s {
		if x.Lifecycle.Is(lif) {
			lis = append(lis, x)
		}
	}

	return lis
}

func (s Slicer) ObjectParent(par objectid.ID) Slicer {
	var lis Slicer

	for _, x := range s {
		if x.Parent == par {
			lis = append(lis, x)
		}
	}

	return lis
}

// --------------------------------------------------------------------- //

func (s Slicer) IDObject(cla objectid.ID) *Object {
	for _, x := range s {
		if cla == x.ID {
			return x
		}
	}

	return nil
}

func (s Slicer) LatestObject() *Object {
	if len(s) == 0 {
		return nil
	}

	var lat *Object
	{
		lat = s[0]
	}

	for _, x := range s[1:] {
		if x.Expiry.After(lat.Expiry) {
			lat = x
		}
	}

	return lat
}
