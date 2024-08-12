package walletstorage

type Slicer []*Object

func (s Slicer) ObjectActive(act bool) Slicer {
	var lis Slicer

	for _, x := range s {
		if act == x.Active {
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
