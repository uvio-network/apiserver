package postvalidator

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/format/labelname"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*poststorage.Object) ([]*poststorage.Object, error) {
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
			inp[i].Labels = generic.Func(inp[i].Labels, labelname.Format)
			inp[i].Text = strings.TrimSpace(inp[i].Text)
			inp[i].Token = strings.TrimSpace(inp[i].Token)
		}

		// Create the existing tree ID, or create a new one. Any proposed claim
		// effectively starts a new tree. So for every post of kind "claim" that has
		// lifecycle "propose" we generate a new tree ID. For all other claims we
		// fetch the post object provided in the parent field of the given claim and
		// take the existing tree ID from there. Note that we are indirectly
		// validating here that the given parent ID does in fact exist.
		if inp[i].Kind == "claim" {
			if inp[i].Lifecycle == "propose" {
				inp[i].Tree = objectid.Random(objectid.Time(now))
			} else {
				var obj []*poststorage.Object
				{
					obj, err = r.sto.Post().SearchPost("", []objectid.ID{inp[i].Parent})
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				if len(obj) != 1 {
					return nil, tracer.Mask(runtime.ExecutionFailedError)
				}

				{
					inp[i].Tree = obj[0].Tree
				}
			}
		}
	}

	return inp, nil
}
