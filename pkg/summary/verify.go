package summary

import "github.com/uvio-network/apiserver/pkg/storage/poststorage"

func Verify(pos *poststorage.Object) bool {
	if len(pos.Summary) < 2 {
		return false
	}

	return pos.Summary[Agreement] != pos.Summary[Disagreement]
}
