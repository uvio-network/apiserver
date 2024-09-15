package walletstorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) UpdateWallet(inp []*Object) error {
	var err error

	for i := range inp {
		err = r.red.Simple().Create().Element(walObj(inp[i].ID), musStr(inp[i]))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
