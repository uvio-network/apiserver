package votestorage

import "github.com/xh3b4sd/objectid"

type Slicer []*Object

func (s Slicer) Claim() []objectid.ID {
	var lis []objectid.ID

	for _, x := range s {
		lis = append(lis, x.Claim)
	}

	return lis
}

func (s Slicer) Owner() []objectid.ID {
	var lis []objectid.ID

	for _, x := range s {
		lis = append(lis, x.Owner)
	}

	return lis
}

func (s Slicer) LifecycleConfirmed() Slicer {
	var lis Slicer

	for _, x := range s {
		if !x.Lifecycle.Pending() {
			lis = append(lis, x)
		}
	}

	return lis
}
