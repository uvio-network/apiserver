package poststorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) DeleteExpiry(inp []*Object) error {
	var err error

	for i := range inp {
		err = r.red.Sorted().Delete().Value(posExp(inp[i].Lifecycle.Data), inp[i].ID.String())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
