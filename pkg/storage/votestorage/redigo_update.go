package votestorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) UpdateVote(inp []*Object) error {
	var err error

	for i := range inp {
		{
			err = inp[i].Verify()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Simple().Create().Element(votObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
