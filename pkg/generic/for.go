package generic

import "github.com/xh3b4sd/objectid"

func For[T string | objectid.ID](max int, fnc func(int) T) []T {
	var key []T

	for i := 0; i < max; i++ {
		key = append(key, fnc(i))
	}

	return key
}
