package poststorage

import (
	"slices"

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

// --------------------------------------------------------------------- //

func (s Slicer) IDClaim(cla objectid.ID) *Object {
	for _, x := range s {
		if cla == x.ID {
			return x
		}
	}

	return nil
}

func (s Slicer) LatestClaim() *Object {
	if len(s) == 0 {
		return nil
	}

	slices.SortFunc(s, func(a, b *Object) int {
		return int(a.Expiry.Unix() - b.Expiry.Unix())
	})

	var las *Object
	{
		las = s[len(s)-1]
	}

	return las
}

func (s Slicer) NextClaim(par objectid.ID) *Object {
	for _, x := range s {
		if x.Parent == par {
			return x
		}
	}

	return nil
}
