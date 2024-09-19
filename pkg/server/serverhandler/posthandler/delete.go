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

func (h *Handler) Delete(ctx context.Context, req *post.DeleteI) (*post.DeleteO, error) {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var ids []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" {
			ids = append(ids, objectid.ID(x.Intern.Id))
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
	// Reconcile the given resources if necessary.
	//

	{
		pos, err = h.rec.Post().DeletePost(pos)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Delete the given resources.
	//

	{
		err = h.sto.Post().DeletePost(pos)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *post.DeleteO
	{
		res = &post.DeleteO{}
	}

	for range pos {
		res.Object = append(res.Object, &post.DeleteO_Object{
			Intern: &post.DeleteO_Object_Intern{
				Status: objectstatus.Deleted,
			},
		})
	}

	return res, nil
}
