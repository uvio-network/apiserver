package generic

import "github.com/uvio-network/apiserver/pkg/object/objectid"

func Func[T string | objectid.ID](lis []T, fnc func(T) string) []string {
	var key []string

	for _, x := range lis {
		key = append(key, fnc(x))
	}

	return key
}
