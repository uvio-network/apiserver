package userstorage

import (
	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/xh3b4sd/redigo/pkg/sorted"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) UpdateReputation(inp []*Object) error {
	var err error

	for i := range inp {
		{
			_, err = r.red.Sorted().Update().Score(storageformat.UserReputation, inp[i].ID.String(), inp[i].Reputation())
			if sorted.IsNotFound(err) {
				{
					err = r.red.Sorted().Create().Score(storageformat.UserReputation, inp[i].ID.String(), inp[i].Reputation())
					if err != nil {
						return tracer.Mask(err)
					}
				}
			} else if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func (r *Redigo) UpdateUser(inp []*Object) error {
	var err error

	for i := range inp {
		{
			err = inp[i].Verify()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Simple().Create().Element(useObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
