package summary

import "github.com/uvio-network/apiserver/pkg/storage/poststorage"

func Verify(res *poststorage.Object) bool {
	if len(res.Summary) < 2 {
		return false
	}

	return res.Summary[Agreement] != res.Summary[Disagreement]
}
