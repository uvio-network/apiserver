package poststorage

import "github.com/xh3b4sd/tracer"

func (r *Redigo) InternCreateExpiry(inp []*Object) error {
	var err error

	for i := range inp {
		if inp[i].Kind == "claim" {
			err = r.red.Sorted().Create().Score(posExp(inp[i].Lifecycle.Data), inp[i].ID.String(), float64(inp[i].Expiry.UnixNano()))
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
