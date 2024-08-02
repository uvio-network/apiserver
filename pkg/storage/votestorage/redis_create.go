package votestorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		// Create the normalized key-value pair for the vote object. With that we
		// can search for vote objects using their IDs.
		{
			err = r.red.Simple().Create().Element(votObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Associate the given vote ID with the given claim ID so that we can search
		// for all votes that have been cast on any given claim. The interesting
		// implication of this particular index is that we can later analyse how
		// users behaved on their own and in relation to one another.
		{
			err = r.red.Sorted().Create().Score(votCla(inp[i].Claim), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Track the user creating this vote as the owner, and make sure that we can
		// find all votes for any given user ID.
		{
			err = r.red.Sorted().Create().Score(votOwn(inp[i].Owner), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
