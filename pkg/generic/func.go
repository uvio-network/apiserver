package generic

import "github.com/xh3b4sd/objectid"

func Func[T string | objectid.ID](lis []T, fnc func(T) string) []string {
	var key []string

	for _, x := range lis {
		key = append(key, fnc(x))
	}

	return key
}
