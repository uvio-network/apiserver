package votestorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) DeleteVote(inp []*Object) error {
	var err error

	for i := range inp {
		{
			err = r.red.Sorted().Delete().Score(votOwnCla(inp[i].Owner, inp[i].Claim), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Sorted().Delete().Score(votOwn(inp[i].Owner), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Sorted().Delete().Score(votCla(inp[i].Claim), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			_, err = r.red.Simple().Delete().Multi(votObj(inp[i].ID))
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
