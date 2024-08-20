package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectstatus"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *post.UpdateI) (*post.UpdateO, error) {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var ids []objectid.ID
	var has []string
	var met []string
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" {
			ids = append(ids, objectid.ID(x.Intern.Id))
		}
		if x.Public != nil && x.Public.Hash != "" {
			has = append(has, x.Public.Hash)
		}
		if x.Public != nil && x.Public.Meta != "" {
			met = append(met, x.Public.Meta)
		}
	}

	//
	// Search for all relevant post objects.
	//

	var pos []*poststorage.Object
	{
		pos, err = h.sto.Post().SearchPost(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Verify the ownership of the given resources.
	//

	for i := range pos {
		if pos[i].Owner != use {
			return nil, tracer.Maskf(runtime.UserNotOwnerError, "%s != %s", pos[i].Owner, use)
		}
	}

	//
	// Update the post hash.
	//

	if len(has) != 0 {
		if len(ids) != len(has) {
			return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(ids), len(has))
		}

		{
			pos, err = h.rec.Post().UpdateHash(pos, has)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	//
	// Update the post meta.
	//

	if len(met) != 0 {
		if len(ids) != len(met) {
			return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(ids), len(met))
		}

		{
			pos, err = h.rec.Post().UpdateMeta(pos, met)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	//
	// Update the given resources.
	//

	{
		err = h.sto.Post().UpdatePost(pos)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *post.UpdateO
	{
		res = &post.UpdateO{}
	}

	for range pos {
		res.Object = append(res.Object, &post.UpdateO_Object{
			Intern: &post.UpdateO_Object_Intern{
				Status: objectstatus.Updated,
			},
		})
	}

	return res, nil
}
