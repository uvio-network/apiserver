package walletstorage

type Slicer []*Object

func (s Slicer) ObjectKind(kin string) Slicer {
	var lis Slicer

	for _, x := range s {
		if x.Kind == kin {
			lis = append(lis, x)
		}
	}

	return lis
}
