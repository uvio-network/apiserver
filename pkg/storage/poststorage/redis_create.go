package poststorage

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/redigo/simple"
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
				var jsn []string
				{
					jsn, err = r.red.Simple().Search().Multi(posObj(inp[i].Parent))
					if simple.IsNotFound(err) {
						return nil, tracer.Maskf(PostObjectNotFoundError, "%v", inp[i].Parent)
					} else if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				if len(jsn) != 1 {
					return nil, tracer.Mask(runtime.ExecutionFailedError)
				}

				var obj *Object
				{
					obj = &Object{}
				}

				{
					err = json.Unmarshal([]byte(jsn[0]), obj)
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				{
					inp[i].Tree = obj.Tree
				}
			}
		}

		// Create the normalized key-value pair for the post object. With that we
		// can search for post objects using their IDs.
		{
			err = r.red.Simple().Create().Element(posObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Add the given claim object ID to the list of label names, so that we can
		// search for claims given any label name.
		for _, x := range inp[i].Labels {
			err = r.red.Sorted().Create().Score(posLab(x), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Associate the given post ID with the given tree ID if the given post kind
		// is "claim". That means we are only managing trees for claims, and not for
		// comments.
		if inp[i].Tree != "" {
			err = r.red.Sorted().Create().Score(posTre(inp[i].Tree), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Track the user creating this post as the owner, and make sure that we
		// can find all posts for any given user ID.
		{
			err = r.red.Sorted().Create().Score(posOwn(inp[i].Owner), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return inp, nil
}
