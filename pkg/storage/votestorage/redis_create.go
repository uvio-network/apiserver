package votestorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*Object) ([]*Object, error) {
	var err error

	for i := range inp {
		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			inp[i].Created = now
			inp[i].ID = objectid.Random(objectid.Time(now))
		}

		// Create the normalized key-value pair for the vote object. With that we
		// can search for vote objects using their IDs.
		{
			err = r.red.Simple().Create().Element(votObj(inp[i].ID), musStr(inp[i]))
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
