package poststorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdatePost(inp []*Object) error {
	var err error

	for i := range inp {
		err = r.red.Simple().Create().Element(posObj(inp[i].ID), musStr(inp[i]))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
