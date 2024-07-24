package poststorage

import (
	"strings"
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
			inp[i].Text = strings.TrimSpace(inp[i].Text)
			inp[i].Token = strings.TrimSpace(inp[i].Token)
		}

		// Create the normalized key-value pair for the post object. With that we
		// can search for post objects using their IDs.
		{
			err = r.red.Simple().Create().Element(posObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
