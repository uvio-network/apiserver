package votehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectstatus"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, req *vote.DeleteI) (*vote.DeleteO, error) {
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
	// Search for all relevant vote objects.
	//

	var vot []*votestorage.Object
	{
		vot, err = h.sto.Vote().SearchVote(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Verify the ownership of the given resources.
	//

	for i := range vot {
		if vot[i].Owner != use {
			return nil, tracer.Maskf(runtime.UserNotOwnerError, "%s != %s", vot[i].Owner, use)
		}
	}

	//
	// Reconcile the given resources if necessary.
	//

	{
		vot, err = h.rec.Vote().DeleteVote(vot)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Delete the given resources.
	//

	{
		err = h.sto.Vote().DeleteVote(vot)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *vote.DeleteO
	{
		res = &vote.DeleteO{}
	}

	for range vot {
		res.Object = append(res.Object, &vote.DeleteO_Object{
			Intern: &vote.DeleteO_Object_Intern{
				Status: objectstatus.Deleted,
			},
		})
	}

	return res, nil
}
