package converter

import (
	"strconv"

	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

func PageToInt(str string) (int64, error) {
	if str == "" {
		return 0, tracer.Mask(runtime.QueryPagingMissingError)
	}

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, tracer.Mask(runtime.QueryPagingInvalidError)
	}

	if num < 0 {
		return 0, tracer.Mask(runtime.QueryPagingNegativeError)
	}

	return num, nil
}
