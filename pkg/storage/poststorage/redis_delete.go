package poststorage

import (
	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteExpiry(inp []*Object) error {
	var err error

	for i := range inp {
		err = r.red.Sorted().Delete().Value(storageformat.PostExpiry, inp[i].ID.String())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
